package provider

import (
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/realotz/mstore/pkg/tusd/filestore"
	tusd "github.com/tus/tusd/pkg/handler"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"sort"
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
	return &localProvider{config: cf, composer: composer, handler: handler}, nil
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
	if p.Exist(ctx, fileName) {
		f, err = os.OpenFile(fileName, os.O_APPEND, 0666)
	} else {
		f, err = os.Create(fileName)
	}
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	return err
}

// 删除文件
func (p *localProvider) Delete(ctx context.Context, fileName string) error {
	return os.Remove(filepath.Join(p.config.Path, fileName))
}

// 打开文件
func (p *localProvider) Open(ctx context.Context, fileName string) ([]byte, error) {
	file, err := os.Open(filepath.Join(p.config.Path, fileName))
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(file)
	_ = file.Close()
	if err != nil {
		return nil, err
	}
	return body, err
}

// 文件列表
func (p *localProvider) List(ctx context.Context, req ListOption) ([]FileInfo, error) {
	var list []FileInfo
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
		switch req.Type{
		case 2:
			if v.IsDir(){
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
			if !v.IsDir(){
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
	switch req.SortFlag {
	case NameDescSort:
		sort.Slice(list, func(i, j int) bool { return list[i].Name > list[j].Name })
		break
	case NameAscSort:
		sort.Slice(list, func(i, j int) bool { return list[i].Name < list[j].Name })
		break
	case TimeDescSort:
		sort.Slice(list, func(i, j int) bool { return list[i].UpdatedAt > list[j].UpdatedAt })
		break
	case TimeAscSort:
		sort.Slice(list, func(i, j int) bool { return list[i].UpdatedAt < list[j].UpdatedAt })
		break
	case SizeDescSort:
		sort.Slice(list, func(i, j int) bool { return list[i].Size > list[j].Size })
		break
	case SizeAscSort:
		sort.Slice(list, func(i, j int) bool { return list[i].Size < list[j].Size })
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
		return nil, err
	}
	return list, nil
}
