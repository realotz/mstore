package storage

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/realotz/mstore/internal/biz/storage/provider"
	"sync"
	"time"
)

type Volume struct {
	Id             uuid.UUID
	Name           string
	ProviderName   string
	ProviderConfig []byte
	Provider       provider.VolumeProvider
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type VolumeRepo interface {
	Store(context.Context, Volume) error
	LoadAll(context.Context) ([]Volume, error)
	Delete(context.Context, string) error
}

// 创建卷管理器
func NewVolumeManager(repo VolumeRepo) (*VolumeManager, error) {
	vs, err := repo.LoadAll(context.Background())
	if err != nil {
		return nil, err
	}
	vm := &VolumeManager{volumes: sync.Map{}, repo: repo}
	for _, v := range vs {
		if v.Provider, err = provider.Parse(v.ProviderName, v.Id.String(), v.ProviderConfig); err != nil {
			return nil, err
		}
		vm.volumes.Store(v.Id.String(), &v)
	}
	return vm, nil
}

// 卷管理器
type VolumeManager struct {
	volumes sync.Map
	repo    VolumeRepo
}

// 获取卷
func (m *VolumeManager) GetVolume(vid string) (*Volume, error) {
	val, ok := m.volumes.Load(vid)
	if !ok {
		return nil, errors.New("VolumeNotFound")
	}
	v, _ := val.(*Volume)
	return v, nil
}

// 获取卷
func (m *VolumeManager) GetVolumeAll(ctx context.Context) ([]*Volume, error) {
	var vs []*Volume
	m.volumes.Range(func(key, value interface{}) bool {
		v, _ := value.(*Volume)
		vs = append(vs, v)
		return true
	})
	return vs, nil
}

// 删除卷
func (m *VolumeManager) DelVolume(ctx context.Context, vid string) error {
	if err := m.repo.Delete(ctx, vid); err != nil {
		return err
	}
	m.volumes.Delete(vid)
	return nil
}

// 创建卷
func (m *VolumeManager) CreateVolume(ctx context.Context, vol Volume) (*Volume, error) {
	if vol.ProviderName == "" {
		return nil, errors.New("volume type is null")
	}
	pro, err := provider.Parse(vol.ProviderName, vol.Id.String(), vol.ProviderConfig)
	if err != nil {
		return nil, err
	}
	vol.Provider = pro
	vol.CreatedAt = time.Now()
	vol.UpdatedAt = time.Now()
	err = m.repo.Store(ctx, vol)
	if err != nil {
		return nil, err
	}
	m.volumes.Store(vol.Id.String(), &vol)
	return &vol, nil
}
