package data

import (
	"context"
	"fmt"
	"git.hxecloud.com/cloudwonder-portal/rms-server/pkg/gplugin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	jsoniter "github.com/json-iterator/go"
	"github.com/patrickmn/go-cache"
	"github.com/realotz/mstore/internal/biz"
	"github.com/realotz/mstore/internal/conf"
	"github.com/realotz/mstore/internal/data/entity"
	"github.com/realotz/mstore/pkg/middleware/transaction"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewUserRepo,
	NewVolumeRepo,
	NewAuthRepo,
)

// Data .
type Data struct {
	db      *gorm.DB
	cache   *cache.Cache
	cfg     *conf.Data
}

// NewData .
func NewData(c *conf.Data,  logger log.Logger) (*Data, error) {
	logHelper := log.NewHelper(logger)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       c.Database.Source,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, //取消外键
		//Logger:                                   mylog.NewGormLogger(logger),
	})
	if err != nil {
		logHelper.Errorf("failed opening connection to mysql: %v", err)
		return nil, err
	}
	db = db.Debug()
	err = autoMigrate(db)
	if err != nil {
		return nil, err
	}
	// 链路中间件
	_ = db.Use(&gplugin.OpentracingPlugin{})
	return &Data{
		db:      db,
		cache:   cache.New(5*time.Minute, 10*time.Minute),
		cfg:     c,
	}, nil
}

// db
func (d Data) DB(ctx context.Context) *gorm.DB {
	return transaction.Begin(ctx, d.db)
}

func (d Data) Transaction(ctx context.Context, fc func(tx *gorm.DB) error) error {
	return fc(transaction.Begin(ctx, d.db))
}

// 缓存数据 目前采用内存缓存
func (d Data) JsonCacheData(_ context.Context, key string, disc interface{}, fc func() (interface{}, error), duration time.Duration) error {
	if data, ok := d.cache.Get(key); ok {
		err := jsoniter.Unmarshal(data.([]byte), disc)
		if err == nil {
			return nil
		}
	}
	data, err := fc()
	if err != nil {
		return err
	}
	jsonBytes, err := jsoniter.Marshal(data)
	if err != nil {
		return err
	}
	d.cache.Set(key, jsonBytes, duration)
	return jsoniter.Unmarshal(jsonBytes, disc)
}

// 关联模式分页查询列表
func (Data) PageFindAssociationList(db *gorm.DB, mod interface{}, list interface{}, total *int64, op biz.ListOption, association string) error {
	if op.Page == 0 {
		op.Page = 1
	}
	*total = db.Model(mod).Association(association).Count()
	if !op.NoPage {
		limit := op.GetPageSize()
		offset := limit * (op.Page - 1)
		db = db.Limit(int(limit)).Offset(int(offset))
	}
	if op.OrderField != "" {
		order := "desc"
		if !op.OrderDesc {
			order = "asc"
		}
		db.Order(fmt.Sprintf("%s %s", op.OrderField, order))
	}
	return db.Model(mod).Association(association).Find(list)
}

// 获取分页列表
func (Data) PageFindList(db *gorm.DB, list interface{}, total *int64, op biz.ListOption, selects ...interface{}) error {
	if op.Page == 0 {
		op.Page = 1
	}
	err := db.Count(total).Error
	if err != nil {
		return err
	}
	if !op.NoPage {
		limit := op.GetPageSize()
		offset := limit * (op.Page - 1)
		db = db.Limit(int(limit)).Offset(int(offset))
	}
	var orders []string
	if op.OrderField != "" {
		order := "desc"
		if !op.OrderDesc {
			order = "asc"
		}
		orders = append(orders, fmt.Sprintf("%s %s", op.OrderField, order))
	}
	for _, v := range op.OrderMore {
		if v.OrderField != "" {
			order := "desc"
			if !v.OrderDesc {
				order = "asc"
			}
			orders = append(orders, fmt.Sprintf("%s %s", v.OrderField, order))
		}
	}
	if len(orders) > 0 {
		db.Order(strings.Join(orders, ","))
	}
	if len(selects) > 0 {
		if len(selects) > 1 {
			db = db.Select(selects[0], selects[1:]...)
		} else {
			db = db.Select(selects[0])
		}
	}
	return db.Find(list).Error
}

// 获取分页列表
func (Data) FindAll(db *gorm.DB, list interface{}, total *int64) error {
	err := db.Count(total).Error
	if err != nil {
		return err
	}
	return db.Find(list).Error
}

// 初始化数据库表
func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		//storage
		&entity.Volume{},

		// users
		&entity.User{},
		&entity.Auth{},
	)
}
