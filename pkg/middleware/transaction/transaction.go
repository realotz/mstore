package transaction

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"gorm.io/gorm"
	"sync"
)

type transactionKey struct{}

type transaction struct {
	sync.RWMutex
	gormDb *gorm.DB
}

func formContext(ctx context.Context) (*transaction, bool) {
	tr, ok := ctx.Value(transactionKey{}).(*transaction)
	return tr, ok
}

func Begin(ctx context.Context, db *gorm.DB) *gorm.DB {
	tr, ok := formContext(ctx)
	if !ok {
		return db.Session(&gorm.Session{NewDB: true, Context: ctx})
	}
	if tr.gormDb == nil {
		tr.gormDb = db.Session(&gorm.Session{NewDB: true, Context: ctx}).Begin()
		return tr.gormDb
	} else {
		return tr.gormDb.Session(&gorm.Session{NewDB: true, Context: ctx})
	}
}

func Start(ctx context.Context) context.Context {
	return context.WithValue(ctx, transactionKey{}, &transaction{})
}

func Rollback(ctx context.Context) {
	tr, ok := formContext(ctx)
	if !ok {
		return
	}
	if tr.gormDb != nil {
		tr.gormDb.Rollback()
	}
}

func Commit(ctx context.Context) {
	tr, ok := formContext(ctx)
	if !ok {
		return
	}
	if tr.gormDb != nil {
		tr.gormDb.Commit()
	}
}

// 全局事务
func Server() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			ctx = Start(ctx)
			reply, err = handler(ctx, req)
			if err != nil {
				Rollback(ctx)
			} else {
				Commit(ctx)
			}
			return reply, err
		}
	}
}
