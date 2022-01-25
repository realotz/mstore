package storage

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"github.com/realotz/mstore/api/errors"
	storageV1 "github.com/realotz/mstore/api/storage/v1"
	"github.com/realotz/mstore/internal/biz/storage/provider"
	"io"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
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
func (s *StorageUseCase) RenameFile(ctx context.Context, id string, path, newPath string, wireType uint32) error {
	volume, err := s.volumeManager.GetVolume(id)
	if err != nil {
		return err
	}
	_, fileName := filepath.Split(path)
	newName := filepath.Join(newPath, fileName)
	if volume.Provider.Exists(ctx, newName) && wireType == 0 {
		return errors.ErrorConflictError("该目录文件名重复")
	}
	if wireType == 1 {
		_ = volume.Provider.Delete(ctx, newName)
	}
	if wireType == 2 {
		n := 0
		newName += " copy"
		sName := ""
		for !volume.Provider.Exists(ctx, newName+sName) {
			n++
			sName = fmt.Sprint(n)
		}
		newName = newName + sName
	}
	return volume.Provider.Rename(ctx, path, newName)
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
		} else {
			volume, err = s.volumeManager.GetVolume(v.Id)
			if err != nil {
				return err
			}
		}
		_, fileName := filepath.Split(v.Path)
		toBasePath := filepath.Join(req.ToPath, fileName)
		if volume.Provider.Exists(ctx, toBasePath) && req.WireType == 0 {
			return errors.ErrorConflictError("该目录文件名重复")
		}
		// 重命名
		if req.WireType == 2 {
			n := 0
			toBasePath += " copy"
			sName := ""
			for volume.Provider.Exists(ctx, toBasePath+sName) {
				n++
				sName = fmt.Sprint(n)
			}
			toBasePath = toBasePath + sName

		}
		// 覆盖
		err = volume.Provider.Walk(v.Path, func(path string, info fs.FileInfo, err error) error {
			if info == nil {
				return nil
			}
			newPath := strings.ReplaceAll(path, v.Path, toBasePath)
			if info.IsDir() {
				fmt.Println(1222, path, v.Path, toBasePath, newPath)
				if !volume.Provider.Exists(ctx, newPath) {
					return volume.Provider.CreateDir(ctx, newPath)
				} else {
					return nil
				}
			} else {
				if req.IsDelete && v.Id == req.ToVolumeId {
					return volume.Provider.Rename(ctx, v.Path, newPath)
				}
				nf, err := newVolume.Provider.Create(ctx, newPath)
				if err != nil {
					return err
				}
				file, err := volume.Provider.Open(ctx, path)
				if err != nil {
					_ = nf.Close()
					return err
				}
				_, err = io.Copy(nf, file)
				if err != nil {
					_ = file.Close()
					_ = nf.Close()
					return err
				}
			}
			return nil
		})
		if err != nil {
			return errors.ErrorConflictError(err.Error())
		}
	}
	return nil
}
