package provider

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/tus/tusd/pkg/filestore"
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
	config  *localConfig
	store   filestore.FileStore
	handler *tusd.UnroutedHandler
}

func NewLocalProvider(data []byte) (*localProvider, error) {
	cf, err := parseConfig(data)
	if err != nil {
		return nil, err
	}
	store := filestore.FileStore{
		Path: cf.Path,
	}
	composer := tusd.NewStoreComposer()
	store.UseIn(composer)
	handler, err := tusd.NewUnroutedHandler(tusd.Config{
		StoreComposer:         composer,
		NotifyCompleteUploads: true,
	})
	if err != nil {
		return nil, err
	}
	return &localProvider{config: cf, store: filestore.FileStore{
		Path: cf.Path,
	}, handler: handler}, nil
}

func (p *localProvider)GetProviderType() string{
	return VolumeTypeLocal
}

// http大文件上传
func (p *localProvider) HttpBigUpload(w http.ResponseWriter,r *http.Request) error {
	p.handler.PostFile(w,r)
	return nil
}

// 小文件上传
func (p *localProvider) Upload(ctx context.Context,fileName string, data []byte) error {
	return nil
}

// 删除文件
func (p *localProvider) Delete(ctx context.Context,fileName string) error {
	return os.Remove(filepath.Join(p.config.Path, fileName))
}

// 打开文件
func (p *localProvider) Open(ctx context.Context,fileName string) ([]byte, error) {
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
func (p *localProvider) List(ctx context.Context,path string, sortFlag FileSort) ([]FileInfo, error) {
	var list []FileInfo
	dirList, err := readDir(filepath.Join(p.config.Path, path))
	if err != nil {
		return nil, err
	}
	for _, v := range dirList {
		list = append(list, FileInfo{
			IsDir:     v.IsDir(),
			Name:      v.Name(),
			Size:      v.Size(),
			Ext:       filepath.Ext(v.Name()),
			UpdatedAt: v.ModTime().Unix(),
		})
	}
	switch sortFlag {
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
