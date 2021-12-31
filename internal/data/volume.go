package data

import (
	"context"
	"github.com/google/uuid"
	"github.com/realotz/mstore/internal/biz/storage"
	"github.com/realotz/mstore/internal/data/entity"
	"github.com/realotz/mstore/pkg/base"
)

func NewVolumeRepo(data *Data) storage.VolumeRepo {
	return &volume{data: data}
}

type volume struct {
	data *Data
}

// 删除卷信息
func(v volume) Delete(ctx context.Context, vid string) error {
	return v.data.DB(ctx).Where("id=?", vid).Delete(&entity.Volume{}).Error
}

// 保存卷信息
func (v volume) Store(ctx context.Context, volume storage.Volume) error {
	return v.data.DB(ctx).Save(&entity.Volume{
		UuidModel:      base.UuidModel{ID: uuid.New()},
		Name:           volume.Name,
		Provider:       volume.ProviderName,
		ProviderConfig: volume.ProviderConfig,
	}).Error
}

// 加载全部存储卷
func (v volume) LoadAll(ctx context.Context) ([]storage.Volume, error) {
	var list []entity.Volume
	err := v.data.DB(ctx).Model(&entity.Volume{}).Find(&list).Error
	if err != nil {
		return nil, err
	}
	var res []storage.Volume
	for _, v := range list {
		res = append(res, storage.Volume{
			Id:             v.ID,
			Name:           v.Name,
			ProviderName:   v.Provider,
			ProviderConfig: v.ProviderConfig,
		})
	}
	return res, nil
}
