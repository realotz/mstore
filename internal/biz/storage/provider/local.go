package provider

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport"
	jsoniter "github.com/json-iterator/go"
	"github.com/realotz/mstore/api/errors"
	"github.com/realotz/mstore/pkg/tusd/filestore"
	tusd "github.com/tus/tusd/pkg/handler"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type localConfig struct {
	Path string `json:"path"`
}

func parseConfig(data []byte) (*localConfig, error) {
	var config localConfig
	err := jsoniter.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

type localProvider struct {
	config   *localConfig
	id       string
	composer *tusd.StoreComposer
	handler  *tusd.UnroutedHandler
}

func NewLocalProvider(id string, data []byte) (*localProvider, error) {
	cf, err := parseConfig(data)
	if err != nil {
		return nil, err
	}
	store := filestore.New(cf.Path)
	composer := tusd.NewStoreComposer()
	store.UseIn(composer)
	handler, err := tusd.NewUnroutedHandler(tusd.Config{
		StoreComposer:         composer,
		NotifyCompleteUploads: true,
		BasePath:              "/api/v1/big-upload/" + id,
	})
	go func() {
		for {
			msg := <-handler.CompleteUploads
			fmt.Println(msg.Upload)
		}
	}()
	if err != nil {
		return nil, err
	}
	return &localProvider{id: id, config: cf, composer: composer, handler: handler}, nil
}

func (p *localProvider) GetProviderType() string {
	return VolumeTypeLocal
}

// http大文件上传
func (p *localProvider) HttpBigUpload(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case "GET":
		p.handler.GetFile(w, r)
		return nil
	case "POST":
		p.handler.PostFile(w, r)
		return nil
	case http.MethodDelete:
		p.handler.DelFile(w, r)
		return nil
	case http.MethodHead:
		p.handler.HeadFile(w, r)
		return nil
	case http.MethodPatch:
		p.handler.PatchFile(w, r)
		return nil
	}
	return nil
}

func (p *localProvider) Exist(ctx context.Context, fileName string) bool {
	var exist = true
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// 小文件上传
func (p *localProvider) Upload(ctx context.Context, fileName string, data []byte) error {
	var f *os.File
	var err error
	fileName = filepath.Join(p.config.Path, fileName)
	if p.Exist(ctx, fileName) {
		f, err = os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	} else {
		f, err = os.Create(fileName)
	}
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	return err
}

// 删除文件
func (p *localProvider) Delete(ctx context.Context, fileName string) error {
	return os.RemoveAll(filepath.Join(p.config.Path, fileName))
}

// 删除文件
func (p *localProvider) Create(ctx context.Context, fileName string) (io.ReadWriteCloser, error) {
	return os.Create(filepath.Join(p.config.Path, fileName))
}

func (p *localProvider) CreateDir(ctx context.Context, path string) error {
	return os.MkdirAll(filepath.Join(p.config.Path, path), 0766)
}

// 打开文件
func (p *localProvider) Open(ctx context.Context, fileName string) (io.ReadWriteCloser, error) {
	return os.Open(filepath.Join(p.config.Path, fileName))
}

func (p *localProvider) GetFileUrl(ctx context.Context, path string) (string, error) {
	host := ""
	if sc, ok := transport.FromServerContext(ctx); ok {
		host = sc.Endpoint()
	}
	return fmt.Sprintf("%s/api/v1/files/%s?p=%s", host, p.id, url.QueryEscape(path)), nil
}

// 遍历
func (p *localProvider) Walk(path string, fn filepath.WalkFunc) error {
	return filepath.Walk(filepath.Join(p.config.Path, path), func(path string, fi os.FileInfo, err error) error {
		return fn(strings.ReplaceAll(path, p.config.Path, ""), fi, err)
	})
}

//
func (p *localProvider) Exists(ctx context.Context, path string) bool {
	path = filepath.Join(p.config.Path, path)
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// 重命名文件
func (p *localProvider) Rename(ctx context.Context, fileName, newName string) error {
	toPath := filepath.Join(p.config.Path, newName)
	return os.Rename(filepath.Join(p.config.Path, fileName), toPath)
}

// 文件列表
func (p *localProvider) List(ctx context.Context, req ListOption) ([]FileInfo, error) {
	var list []FileInfo
	req.Path = filepath.Join(filepath.Split(req.Path))
	dirList, err := readDir(filepath.Join(p.config.Path, req.Path))
	if err != nil {
		return nil, err
	}
	for _, v := range dirList {
		if !req.HideFile {
			if string(v.Name()[0]) == "." {
				continue
			}
		}
		switch req.Type {
		case 2:
			if v.IsDir() {
				list = append(list, FileInfo{
					IsDir:     v.IsDir(),
					Path:      req.Path,
					Name:      v.Name(),
					Size:      v.Size(),
					Ext:       filepath.Ext(v.Name()),
					UpdatedAt: v.ModTime().Unix(),
				})
			}
			break
		case 3:
			if !v.IsDir() {
				list = append(list, FileInfo{
					IsDir:     v.IsDir(),
					Path:      req.Path,
					Name:      v.Name(),
					Size:      v.Size(),
					Ext:       filepath.Ext(v.Name()),
					UpdatedAt: v.ModTime().Unix(),
				})
			}
			break
		default:
			list = append(list, FileInfo{
				IsDir:     v.IsDir(),
				Path:      req.Path,
				Name:      v.Name(),
				Size:      v.Size(),
				Ext:       filepath.Ext(v.Name()),
				UpdatedAt: v.ModTime().Unix(),
			})
		}
	}
	switch req.Sort {
	case NameSort:
		if req.SortDesc {
			sort.Slice(list, func(i, j int) bool { return list[i].Name > list[j].Name })
		} else {
			sort.Slice(list, func(i, j int) bool { return list[i].Name < list[j].Name })
		}
		break
	case TimeSort:
		if req.SortDesc {
			sort.Slice(list, func(i, j int) bool { return list[i].UpdatedAt > list[j].UpdatedAt })
		} else {
			sort.Slice(list, func(i, j int) bool { return list[i].UpdatedAt < list[j].UpdatedAt })
		}
		break
	case SizeSort:
		if req.SortDesc {
			sort.Slice(list, func(i, j int) bool { return list[i].Size > list[j].Size })
		} else {
			sort.Slice(list, func(i, j int) bool { return list[i].Size < list[j].Size })
		}
		break
	case ExtSort:
		if req.SortDesc {
			sort.Slice(list, func(i, j int) bool { return list[i].Ext > list[j].Ext })
		} else {
			sort.Slice(list, func(i, j int) bool { return list[i].Ext < list[j].Ext })
		}
		break
	}
	return list, nil
}

// 读取目录
func readDir(dirname string) ([]fs.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	_ = f.Close()
	if err != nil {
		return nil, errors.ErrorBusinessError("找不到该目录")
	}
	return list, nil
}
