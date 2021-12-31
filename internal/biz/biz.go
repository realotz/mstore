package biz

import (
	"github.com/google/wire"
	v1 "github.com/realotz/mstore/api/core/v1"
	"github.com/realotz/mstore/internal/biz/storage"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	storage.ProviderSet,
	NewAuthUseCase,
	NewUserUseCase,
)

type OrderBy struct {
	// 排序字段
	OrderField string
	// 是否倒序 true desc false asc
	OrderDesc bool
}

// 基础列表查询选项
type ListOption struct {
	// 排序字段
	OrderField string
	// 是否倒序 true desc false asc
	OrderDesc bool
	// 更多排序
	OrderMore []OrderBy
	// 页码
	Page uint32
	// 分页大小
	PageSize uint32
	// 不分页
	NoPage bool
}

func NewListOption(option *v1.ListOption) ListOption {
	if option==nil{
		return ListOption{
			Page:       1,
			PageSize:   50,
		}
	}
	return ListOption{
		OrderField: option.OrderField,
		OrderDesc:  option.OrderDesc,
		Page:       option.Page,
		PageSize:   option.PageSize,
	}
}

// 获取分页大小
func (op *ListOption) GetPageSize() uint32 {
	if op.PageSize == 0 {
		return 200
	}
	return op.PageSize
}