package storage

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/realotz/mstore/api/core/v1"
	"github.com/realotz/mstore/api/errors"
	storageV1 "github.com/realotz/mstore/api/storage/v1"
	"github.com/realotz/mstore/internal/biz/storage"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	vol,err := s.uc.CreateVolume(ctx,storage.Volume{
		Name:           req.Name,
		ProviderName:   req.Provider,
		ProviderConfig: []byte(req.ProviderConfig),
	})
	if err != nil {
		return nil, errors.ErrorBusinessError(err.Error())
	}
	return &storageV1.Volume{
		Name:           vol.Name,
		Provider:       vol.ProviderName,
		ProviderConfig: string(vol.ProviderConfig),
		CreatedAt:      timestamppb.New(vol.CreatedAt),
		UpdatedAt:      timestamppb.New(vol.UpdatedAt),
	},nil
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
			Name:           v.Name,
			Provider:       v.ProviderName,
			ProviderConfig: string(v.ProviderConfig),
			CreatedAt:      timestamppb.New(v.CreatedAt),
			UpdatedAt:      timestamppb.New(v.UpdatedAt),
		})
	}
	return resp, nil
}

// 大文件上传
func (s *VolumeService) BigFileUpload(ctx http.Context) error {
	vid := ctx.Vars().Get("id")
	if vid == "" {
		return errors.ErrorParamsError("volume id is empty!")
	}
	err := s.uc.BigFileUpload(vid, ctx.Response(), ctx.Request())
	if err != nil {
		return errors.ErrorBusinessError(err.Error())
	}
	return nil
}
