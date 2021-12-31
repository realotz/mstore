package entity

import (
	"github.com/realotz/mstore/internal/biz"
	"github.com/realotz/mstore/pkg/base"
	"time"
)

// 用户信息
type User struct {
	base.Model
	Name   string `gorm:"name;comment:名称" json:"name"`
	Phone  string `gorm:"phone;comment:手机" json:"phone"`
	Email  string `gorm:"email;comment:邮箱" json:"email"`
	AuthId uint32 `gorm:"auth_id;comment:认证账号id" json:"auth_id"`
	Auth   *Auth
}

type Users []User

func (ms Users) Convert() []*biz.User {
	var list []*biz.User
	for _, v := range ms {
		list = append(list, v.Convert())
	}
	return list
}

func (u *User) Convert() *biz.User {
	m := &biz.User{
		Id:        u.ID,
		Name:      u.Name,
		Phone:     u.Phone,
		Email:     u.Email,
		AuthId:    u.AuthId,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
	if u.Auth!=nil{
		m.Auth = u.Auth.Convert()
	}
	return m
}

// 认证相关信息
type Auth struct {
	base.Model
	Account       string     `gorm:"account;comment:账号;uniqueIndex:account" json:"account"`
	Salt          string     `gorm:"salt;comment:加密盐" json:"salt"`
	Password      string     `gorm:"password;comment:密码" json:"password"`
	Role          string     `gorm:"role;comment:角色" json:"role"`
	LastIp        string     `gorm:"last_ip;comment:最后登陆ip" json:"last_ip"`
	LastLoginTime *time.Time `gorm:"last_login_time;comment:最后登陆时间" json:"last_login_time"`
}

type Auths []Auth

func (ms Auths) Convert() []*biz.Auth {
	var list []*biz.Auth
	for _, v := range ms {
		list = append(list, v.Convert())
	}
	return list
}

func (u *Auth) Convert() *biz.Auth {
	return &biz.Auth{
		Id:            u.ID,
		Account:       u.Account,
		Password:      u.Password,
		Salt:          u.Salt,
		Role:          u.Role,
		LastIp:        u.LastIp,
		LastLoginTime: u.LastLoginTime,
	}
}
