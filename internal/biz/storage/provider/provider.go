package provider

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

const (
	VolumeTypeLocal = "local"
	VolumeTypeS3    = "s3"
	VolumeTypeOss   = "oss"
	VolumeTypeCos   = "cos"
)

type FileSort string

const (
	NoSort   = FileSort("")
	NameSort = FileSort("name")
	ExtSort  = FileSort("ext")
	TimeSort = FileSort("updated_at")
	SizeSort = FileSort("size")
)

type ListOption struct {
	Path     string
	HideFile bool
	Type     uint32
	Sort     FileSort
	SortDesc bool
}

type VolumeProvider interface {
	// 供应商类型
	GetProviderType() string
	// 大文件上传
	HttpBigUpload(w http.ResponseWriter, r *http.Request) error
	// 上传文件
	Upload(ctx context.Context, fileName string, data []byte) error
	// 删除文件
	Delete(ctx context.Context, fileName string) error
	// 打开文件
	Open(ctx context.Context, name string) (io.ReadWriteCloser, error)
	// 创建文件
	Create(ctx context.Context, name string) (io.ReadWriteCloser, error)
	// 文件列表
	List(ctx context.Context, req ListOption) ([]FileInfo, error)
	// 重命名
	Rename(context.Context, string, string, bool) error
}

type FileInfo struct {
	IsDir     bool   `json:"is_dir"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	Size      int64  `json:"size"`
	Ext       string `json:"ext"`
	UpdatedAt int64  `json:"updated_at"`
}

func Parse(t string, id string, config []byte) (VolumeProvider, error) {
	switch t {
	case VolumeTypeLocal:
		return NewLocalProvider(id, config)
	case VolumeTypeS3:
		return NewLocalProvider(id, config)
	case VolumeTypeOss:
		return NewLocalProvider(id, config)
	case VolumeTypeCos:
		return NewLocalProvider(id, config)
	default:
		return nil, fmt.Errorf("the provider [%s] is not implemented", t)
	}
}
