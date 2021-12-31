package base

import (
	"github.com/google/uuid"
	"time"
)

type Model struct {
	ID        uint32 `gorm:"primarykey;comment:主键id"`
	CreatedAt time.Time `gorm:"created_at;comment:创建时间"`
	UpdatedAt time.Time `gorm:"updated_at;comment:更新时间"`
}

type UuidModel struct {
	ID        uuid.UUID `gorm:"primarykey;comment:主键id"`
	CreatedAt time.Time `gorm:"created_at;comment:创建时间"`
	UpdatedAt time.Time `gorm:"updated_at;comment:更新时间"`
}

