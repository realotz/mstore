package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/realotz/mstore/internal/biz"
	"github.com/realotz/mstore/internal/data/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// 用户与认证repo接口实现
type user struct {
	data *Data
}

func NewUserRepo(data *Data) biz.UserRepo {
	return &user{data: data}
}

// 用户查询条件
func UserOptionWhere(op biz.ListUserOption) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if op.Id != 0 {
			db.Where("id=?", op.Id)
		}
		if op.Keyword != "" {
			keyword := fmt.Sprintf("%%%s%%", op.Keyword)
			db.Where("name like ? or phone like ? or email like ? or id=?", keyword, keyword, keyword, op.Keyword)
		}
		if op.Role != "" {
			rowSelect := db.Session(&gorm.Session{NewDB: true}).Model(&entity.Auth{}).Select("id").Where("role=?", op.Role)
			db.Where("auth_id in (?)", rowSelect)
		}
		if op.Account != "" {
			rowSelect := db.Session(&gorm.Session{NewDB: true}).Model(&entity.Auth{}).Select("id").Where("account=?", op.Account)
			db.Where("auth_id in (?) or phone=? or email = ?", rowSelect, op.Account, op.Account)
		}
		return db
	}
}

// 查询用户列表
func (u *user) ListUser(ctx context.Context, op biz.ListUserOption) ([]*biz.User, int64, error) {
	db := u.data.DB(ctx).Model(&entity.User{}).Scopes(UserOptionWhere(op))
	var list entity.Users
	var total int64
	db.Preload(clause.Associations)
	err := u.data.PageFindList(db, &list, &total, op.ListOption)
	if err != nil {
		return nil, 0, err
	}
	return list.Convert(), total, nil
}

// 获取单个用户
func (u *user) GetUser(ctx context.Context, op biz.ListUserOption) (*biz.User, error) {
	db := u.data.DB(ctx).Model(&entity.User{}).Scopes(UserOptionWhere(op))
	var row entity.User
	db.Preload(clause.Associations)
	err :=db.First(&row).Error
	if err != nil {
		return nil, err
	}
	return row.Convert(), nil
}

// 创建用户
func (u *user) CreateUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	err := u.data.Transaction(ctx, func(tx *gorm.DB) error {
		if user.Auth == nil {
			return errors.New("auth info is null")
		}
		err := tx.Model(&entity.Auth{}).Where("account", user.Auth.Account).First(&entity.Auth{}).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("该账号已存在")
		}
		auth := &entity.Auth{
			Account:  user.Auth.Account,
			Salt:     user.Auth.Salt,
			Password: user.Auth.Password,
			Role:     user.Auth.Role,
		}
		if err = tx.Model(&auth).Create(&auth).Error; err != nil {
			return err
		}
		mod := &entity.User{
			Name:   user.Name,
			Phone:  user.Phone,
			Email:  user.Email,
			AuthId: auth.ID,
		}
		if err := tx.Model(mod).Create(mod).Error; err != nil {
			return err
		}
		mod.Auth = auth
		user = mod.Convert()
		return nil
	})
	return user, err
}

// 更新认证信息
func (u *user) UpdateAuth(ctx context.Context, authId uint32, auth *biz.Auth) (*biz.Auth, error) {
	err := u.data.Transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.Model(&entity.Auth{}).Where("id=?", authId).Updates(&entity.Auth{
			Salt:          auth.Salt,
			Password:      auth.Password,
			Role:          auth.Role,
			LastIp:        auth.LastIp,
			LastLoginTime: auth.LastLoginTime,
		}).Error; err != nil {
			return err
		}
		return nil
	})
	return auth, err
}

// 修改用户信息
func (u *user) UpdateUser(ctx context.Context, id uint32, user *biz.User) (*biz.User, error) {
	err := u.data.Transaction(ctx, func(tx *gorm.DB) error {
		mod := &entity.User{
			Name:  user.Name,
			Phone: user.Phone,
			Email: user.Email,
		}
		if err := tx.Model(mod).Where("id=?", id).Updates(mod).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return u.GetUser(ctx, biz.ListUserOption{Id: id})
}

// 删除用户
func (u *user) DeleteUser(ctx context.Context, uid uint32) error {
	return u.data.Transaction(ctx, func(tx *gorm.DB) error {
		var mod entity.User
		if err := tx.First(&mod, "id=?", uid).Error; err != nil {
			return err
		}
		err := tx.Where("id=?", mod.AuthId).Delete(&entity.Auth{}).Error
		if err != nil {
			return err
		}
		return tx.Delete(&mod).Error
	})
}
