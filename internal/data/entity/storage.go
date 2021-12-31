package entity

import (
	"github.com/realotz/mstore/pkg/base"
	"gorm.io/datatypes"
)

type Volume struct {
	base.UuidModel
	Name           string         `gorm:"name;comment:卷名"`
	Provider       string         `gorm:"provider;comment:提供者"`
	ProviderConfig datatypes.JSON `gorm:"provider_config;comment:供应配置"`
}
