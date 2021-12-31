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

type FileSort uint

const (
	NoSort       = FileSort(0)
	NameDescSort = FileSort(1)
	NameAscSort  = FileSort(2)
	TimeDescSort = FileSort(3)
	TimeAscSort  = FileSort(4)
	SizeDescSort = FileSort(5)
	SizeAscSort  = FileSort(6)
)

type VolumeProvider interface {
	GetProviderType() string
	HttpBigUpload(w http.ResponseWriter, r *http.Request) error
	Upload(ctx context.Context, fileName string, data []byte) error
	Delete(ctx context.Context, fileName string) error
	Open(ctx context.Context, fileName string) ([]byte, error)
	List(ctx context.Context, path string, sortFlag FileSort) ([]FileInfo, error)
}

type FileInfo struct {
	IsDir     bool   `json:"is_dir"`
	Name      string `json:"name"`
	Size      int64  `json:"size"`
	Ext       string `json:"ext"`
	UpdatedAt int64  `json:"updated_at"`
}

func Parse(t string, config []byte) (VolumeProvider, error) {
	switch t {
	case VolumeTypeLocal:
		return NewLocalProvider(config)
	case VolumeTypeS3:
		return NewLocalProvider(config)
	case VolumeTypeOss:
		return NewLocalProvider(config)
	case VolumeTypeCos:
		return NewLocalProvider(config)
	default:
		return nil, fmt.Errorf("the provider [%s] is not implemented", t)
	}
}
