package biz

import (
	"context"
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	captcha "github.com/mojocn/base64Captcha"
	"github.com/realotz/mstore/internal/conf"
	"github.com/realotz/mstore/pkg/myctx"
	"github.com/realotz/mstore/pkg/myjwt"
	"github.com/realotz/mstore/pkg/utils"
	"strings"
	"time"
)

type AuthRepo interface {
	// 获取密码错误统计
	CountPasswdError(ctx context.Context, account string) int
	// 添加密码错误统计
	AddPasswdError(ctx context.Context, account string)
	// 清除统计
	DelPasswdError(ctx context.Context, account string)
}

// 认证信息实体
type Auth struct {
	Id            uint32
	Account       string
	Password      string
	Salt          string
	Role          string
	LastIp        string
	LastLoginTime *time.Time
}

// 认证信息参数检查
func (u *Auth) CreateCheck() error {
	if u.Account == "" {
		return errors.New("账号不能为空")
	}
	if u.Password == "" {
		return errors.New("密码不能为空")
	}
	if u.Role == "" {
		return errors.New("请分配一个角色给用户")
	}
	if u.Salt == "" {
		u.SetPassword(u.Password)
	}
	return nil
}

// 全局密码salt
const groupSalt = ""

// 密码加盐
func encryptPassword(passwd, salt string) string {
	hmacEnt := hmac.New(md5.New, []byte(fmt.Sprintf("%s%s", groupSalt, salt)))
	hmacEnt.Write([]byte(passwd))
	return hex.EncodeToString(hmacEnt.Sum([]byte(nil)))
}

// 检查认证
func (a *Auth) CheckAuth(passwd string) error {
	if passwd == "" {
		return errors.New("密码不能为空")
	}
	if a.Password == "" {
		return errors.New("认证信息查询不到密码")
	}
	if a.Salt == "" {
		return errors.New("认证信息查询不到密钥")
	}
	if a.Password == encryptPassword(passwd, a.Salt) {
		return nil
	} else {
		return fmt.Errorf("账号或密码错误")
	}
}

// 设置密码
func (a *Auth) SetPassword(password string) {
	a.Salt = utils.RandomString(6)
	a.Password = encryptPassword(password, a.Salt)
}

// 认证用例
type AuthUseCase struct {
	tk           *myjwt.PrivateToken
	user         UserRepo
	issues       string
	captchaStore captcha.Store
	driver       *captcha.DriverString
	auth         AuthRepo
}

// wire 创建认证用例接口
func NewAuthUseCase(cfg *conf.Data, user UserRepo, auth AuthRepo) (*AuthUseCase, error) {
	tk, err := myjwt.New(cfg.Auth.Key, cfg.Auth.Cert)
	if err != nil {
		return nil, err
	}
	return &AuthUseCase{
		tk:           tk,
		issues:       "mstore",
		user:         user,
		driver:       utils.NewCaptchaDriver(),
		captchaStore: captcha.DefaultMemStore,
		auth:         auth,
	}, nil
}

// 获取图形验证码
func (uc *AuthUseCase) Captcha(_ context.Context, uuid string) (string, error) {
	d := captcha.NewCaptcha(uc.driver, uc.captchaStore)
	_, content, answer := d.Driver.GenerateIdQuestionAnswer()
	item, _ := d.Driver.DrawCaptcha(content)
	if err := d.Store.Set(uuid, answer); err != nil {
		return "", err
	}
	return item.EncodeB64string(), nil
}

// 图形验证码验证
func (uc *AuthUseCase) ValidateCaptcha(_ context.Context, uuid, code string) bool {
	return uc.captchaStore.Verify(uuid, code, false)
}

// 风险检查
func (uc *AuthUseCase) CheckRisk(ctx context.Context, account string) bool {
	if uc.auth.CountPasswdError(ctx, account) > 4 {
		return true
	}
	return false
}

// 创建token
func (uc *AuthUseCase) CreateUserToken(up myctx.JwtProfile, ut string) (string, error) {
	return uc.tk.Encode(uc.issues, up, ut)
}

// 解析token
func (uc *AuthUseCase) ParseUserToken(ctx context.Context, tokenString string) (*myctx.JwtCustomClaims, error) {
	tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")
	cc, err := uc.tk.Decode(tokenString)
	if err != nil {
		return nil, err
	}
	_, err = uc.user.GetUser(ctx, ListUserOption{
		Id: cc.ID,
	})
	if err != nil {
		return nil, err
	}
	return cc, nil
}

// 刷新token
func (uc *AuthUseCase) RefreshUserToken(ctx context.Context, tokenString string) (string, error) {
	claims, err := uc.ParseUserToken(ctx, tokenString)
	if err != nil {
		return "", err
	}
	return uc.tk.Encode(uc.issues, claims.JwtProfile, claims.UserType)
}

// 查询用户
func (uc *AuthUseCase) GetUser(ctx context.Context, uid uint32) (*User, error) {
	return uc.user.GetUser(ctx, ListUserOption{
		Id: uid,
	})
}

//密码登录
func (uc *AuthUseCase) LoginForPassword(ctx context.Context, account, passwd string) (*User, error) {
	user, err := uc.user.GetUser(ctx, ListUserOption{
		Account: account,
	})
	if err != nil {
		return nil, fmt.Errorf("账号或密码错误")
	}
	if err = user.Auth.CheckAuth(passwd); err == nil {
		uc.auth.DelPasswdError(ctx, account)
		now := time.Now()
		user.Auth.LastLoginTime = &now
		user.Auth.LastIp = utils.GetClientIp(ctx)
		user.Auth, _ = uc.user.UpdateAuth(ctx, user.AuthId, user.Auth)
		return user, nil
	} else {
		uc.auth.AddPasswdError(ctx, account)
		return nil, err
	}
}

//修改密码
func (uc *AuthUseCase) UpdatePassword(ctx context.Context, account, passwd, oldPasswd string) error {
	user, err := uc.user.GetUser(ctx, ListUserOption{
		Account: account,
	})
	if err != nil {
		return fmt.Errorf("账号不存在")
	}

	if err = user.Auth.CheckAuth(oldPasswd); err != nil {
		return err
	}
	//修改密码
	user.Auth.SetPassword(passwd)
	_, err = uc.user.UpdateAuth(ctx, user.AuthId, user.Auth)
	return err
}
