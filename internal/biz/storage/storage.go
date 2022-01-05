package storage

import (
	"context"
	"github.com/google/wire"
	"github.com/realotz/mstore/internal/biz/storage/provider"
	"net/http"
)

var ProviderSet = wire.NewSet(NewVolumeManager, NewStorageUseCase)

type StorageUseCase struct {
	volumeManager *VolumeManager
}

func NewStorageUseCase(volumeManager *VolumeManager) *StorageUseCase {
	return &StorageUseCase{
		volumeManager: volumeManager,
	}
}

// 大文件上传
func (s *StorageUseCase) BigFileUpload(vid string, w http.ResponseWriter, r *http.Request) error {
	volume, err := s.volumeManager.GetVolume(vid)
	if err != nil {
		return err
	}
	return volume.Provider.HttpBigUpload(w, r)
}

// 存储卷列表
func (s *StorageUseCase) ListVolume(ctx context.Context) ([]*Volume, error) {
	return s.volumeManager.GetVolumeAll(ctx)
}

// 删除存储卷
func (s *StorageUseCase) DelVolume(ctx context.Context, vid string) error {
	return s.volumeManager.DelVolume(ctx, vid)
}

// 创建存储卷
func (s *StorageUseCase) CreateVolume(ctx context.Context, vol Volume) (*Volume, error) {
	return s.volumeManager.CreateVolume(ctx, vol)
}

// 文件列表
func (s *StorageUseCase) ListFile(ctx context.Context,id string, op provider.ListOption) ([]provider.FileInfo, error) {
	volume, err := s.volumeManager.GetVolume(id)
	if err != nil {
		return nil, err
	}
	return volume.Provider.List(ctx, op)
}
