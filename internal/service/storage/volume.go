package storage

import (
	"context"
	v1 "github.com/realotz/mstore/api/core/v1"
	"github.com/realotz/mstore/api/errors"
	storageV1 "github.com/realotz/mstore/api/storage/v1"
	"github.com/realotz/mstore/internal/biz/storage"
	"github.com/realotz/mstore/internal/biz/storage/provider"
	"google.golang.org/protobuf/types/known/timestamppb"
	http2 "net/http"
	"strings"
)

func NewVolumeService(uc *storage.StorageUseCase) *VolumeService {
	return &VolumeService{
		uc: uc,
	}
}

type VolumeService struct {
	storageV1.UnimplementedVolumeServiceServer
	uc *storage.StorageUseCase
}

// 创建存储卷
func (s *VolumeService) CreateVolume(ctx context.Context, req *storageV1.CreateVolumeReq) (*storageV1.Volume, error) {
	vol, err := s.uc.CreateVolume(ctx, storage.Volume{
		Name:           req.Name,
		ProviderName:   req.Provider,
		ProviderConfig: []byte(req.ProviderConfig),
	})
	if err != nil {
		return nil, errors.ErrorBusinessError(err.Error())
	}
	return &storageV1.Volume{
		Id:             vol.Id.String(),
		Name:           vol.Name,
		Provider:       vol.ProviderName,
		ProviderConfig: string(vol.ProviderConfig),
		CreatedAt:      timestamppb.New(vol.CreatedAt),
		UpdatedAt:      timestamppb.New(vol.UpdatedAt),
	}, nil
}

func (s *VolumeService) DeleteVolume(ctx context.Context, req *storageV1.DeleteVolumeReq) (*v1.Empty, error) {
	err := s.uc.DelVolume(ctx, req.Id)
	if err != nil {
		return nil, errors.ErrorBusinessError(err.Error())
	}
	return &v1.Empty{}, nil
}

// 存储卷列表
func (s *VolumeService) ListVolume(ctx context.Context, req *storageV1.ListVolumeReq) (*storageV1.ListVolumeReply, error) {
	vs, err := s.uc.ListVolume(ctx)
	if err != nil {
		return nil, errors.ErrorBusinessError(err.Error())
	}
	var resp = &storageV1.ListVolumeReply{
		List:  make([]*storageV1.Volume, 0, len(vs)),
		Total: int64(len(vs)),
	}
	for _, v := range vs {
		resp.List = append(resp.List, &storageV1.Volume{
			Id:             v.Id.String(),
			Name:           v.Name,
			Provider:       v.ProviderName,
			ProviderConfig: string(v.ProviderConfig),
			CreatedAt:      timestamppb.New(v.CreatedAt),
			UpdatedAt:      timestamppb.New(v.UpdatedAt),
		})
	}
	return resp, nil
}

// 文件列表
func (s *VolumeService) ListFile(ctx context.Context, req *storageV1.ListFileReq) (*storageV1.ListFileReply, error) {
	op := provider.ListOption{
		Path:     req.Path,
		HideFile: false,
		Type:     req.Type,
	}
	if req.Option != nil {
		op.Sort = provider.FileSort(req.Option.OrderField)
		op.SortDesc = req.Option.OrderDesc
	}
	list, err := s.uc.ListFile(ctx, req.Id, op)
	if err != nil {
		return nil, errors.ErrorBusinessError(err.Error())
	}
	var resp = &storageV1.ListFileReply{
		List:  make([]*storageV1.File, 0, len(list)),
		Total: int64(len(list)),
	}
	for _, v := range list {
		resp.List = append(resp.List, &storageV1.File{
			Name:      v.Name,
			Size:      v.Size,
			Path:      v.Path,
			Ext:       v.Ext,
			IsDir:     v.IsDir,
			VolumeId:  req.Id,
			UpdatedAt: v.UpdatedAt,
		})
	}
	return resp, nil
}

func (s *VolumeService) ServeHTTP(w http2.ResponseWriter, r *http2.Request) {
	paths := strings.Split(r.URL.String(), "/")
	if len(paths) < 5 {
		w.WriteHeader(400)
		_, _ = w.Write([]byte("Invalid volume id"))
		return
	}
	_ = s.uc.BigFileUpload(paths[4], w, r)
}
