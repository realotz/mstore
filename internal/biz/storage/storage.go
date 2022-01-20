package storage

import (
	"context"
	"github.com/google/wire"
	storageV1 "github.com/realotz/mstore/api/storage/v1"
	"github.com/realotz/mstore/internal/biz/storage/provider"
	"io"
	"net/http"
	"path/filepath"
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
func (s *StorageUseCase) ListFile(ctx context.Context, id string, op provider.ListOption) ([]provider.FileInfo, error) {
	volume, err := s.volumeManager.GetVolume(id)
	if err != nil {
		return nil, err
	}
	return volume.Provider.List(ctx, op)
}

// 删除文件
func (s *StorageUseCase) DelFile(ctx context.Context, id string, path string) error {
	volume, err := s.volumeManager.GetVolume(id)
	if err != nil {
		return err
	}
	return volume.Provider.Delete(ctx, path)
}

// 重命名文件
func (s *StorageUseCase) RenameFile(ctx context.Context, id string, path, newPath string) error {
	volume, err := s.volumeManager.GetVolume(id)
	if err != nil {
		return err
	}
	return volume.Provider.Rename(ctx, path, newPath)
}

// 复制/移动文件
func (s *StorageUseCase) MoveFile(ctx context.Context, req *storageV1.MoveCopyFileReq) error {
	newVolume, err := s.volumeManager.GetVolume(req.ToVolumeId)
	if err != nil {
		return err
	}
	for _, v := range req.Files {
		var volume *Volume
		if v.Id == req.ToVolumeId {
			volume = newVolume
		}
		_, fileName := filepath.Split(v.Path)
		if req.IsDelete && v.Id == req.ToVolumeId {
			if err = volume.Provider.Rename(ctx, v.Path, filepath.Join(req.ToPath, fileName)); err != nil {
				return err
			}
		} else {
			volume, err = s.volumeManager.GetVolume(v.Id)
			if err != nil {
				return err
			}
			file, err := volume.Provider.Open(ctx, v.Path)
			if err != nil {
				return err
			}
			nf, err := newVolume.Provider.Create(ctx, filepath.Join(req.ToPath, fileName))
			if err != nil {
				_ = file.Close()
				return err
			}
			_, err = io.Copy(nf, file)
			if err != nil {
				_ = file.Close()
				_ = nf.Close()
				return err
			}
		}
	}
	return nil
}
