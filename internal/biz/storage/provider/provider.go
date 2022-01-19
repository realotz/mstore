package provider

import (
	"context"
	"fmt"
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
	ExtSort = FileSort("ext")
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
	GetProviderType() string
	HttpBigUpload(w http.ResponseWriter, r *http.Request) error
	Upload(ctx context.Context, fileName string, data []byte) error
	Delete(ctx context.Context, fileName string) error
	Open(ctx context.Context, fileName string) ([]byte, error)
	List(ctx context.Context, req ListOption) ([]FileInfo, error)
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
