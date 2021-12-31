package biz

import (
	"context"
	"errors"
	"time"
)

type UserRepo interface {
	// 获取用户
		GetUser(ctx context.Context, op ListUserOption) (*User, error)
	// 用户列表
	ListUser(ctx context.Context, op ListUserOption) ([]*User, int64, error)
	// 存储新用户 创建新的认证信息
	CreateUser(ctx context.Context, user *User) (*User, error)
	// 更新认证信息
	UpdateAuth(ctx context.Context, id uint32, auth *Auth) (*Auth, error)
	// 更新用户
	UpdateUser(ctx context.Context, id uint32, user *User) (*User, error)
	// 删除用户 同时删除认证信息
	DeleteUser(ctx context.Context, uid uint32) error
}

// 用户查询选项
type ListUserOption struct {
	Id      uint32
	Keyword string
	Role    string
	Account string
	ListOption
}

// 用户结构
type User struct {
	Id    uint32
	Name  string
	Phone string
	Email string
	// 认证相关信息
	AuthId    uint32
	Auth      *Auth
	CreatedAt time.Time
	UpdatedAt time.Time
}

// 检查用户信息
func (u *User) CreateCheck() error {
	if u.Auth == nil {
		return errors.New("帐户认证信息不能为空")
	}
	if u.Name == "" {
		return errors.New("用户信息不能为空")
	}
	return u.Auth.CreateCheck()
}

type UserUseCase struct {
	repo UserRepo
}

func NewUserUseCase(userRepo UserRepo) *UserUseCase {
	return &UserUseCase{repo: userRepo}
}

// 获取用户信息
func (uc *UserUseCase) GetUser(ctx context.Context, uid uint32) (*User, error) {
	return uc.repo.GetUser(ctx, ListUserOption{
		Id: uid,
	})
}

// 用户列表
func (uc *UserUseCase) ListUser(ctx context.Context, op ListUserOption) ([]*User, int64, error) {
	return uc.repo.ListUser(ctx, op)
}

// 删除用户
func (uc *UserUseCase) DelUser(ctx context.Context, uid uint32) error {
	return uc.repo.DeleteUser(ctx, uid)
}

// 创建用户
func (uc *UserUseCase) CreateUser(ctx context.Context, user *User) (*User, error) {
	if err := user.CreateCheck(); err != nil {
		return nil, err
	}
	return uc.repo.CreateUser(ctx, user)
}

// 修改用户信息
func (uc *UserUseCase) UpdateUser(ctx context.Context, uid uint32, user *User) (*User, error) {
	oldUser, err := uc.GetUser(ctx, uid)
	if err != nil {
		return nil, err
	}
	if user.Auth != nil {
		user.Auth, err = uc.repo.UpdateAuth(ctx, oldUser.AuthId, user.Auth)
		if err != nil {
			return nil, err
		}
	}
	return uc.repo.UpdateUser(ctx, uid, user)
}
