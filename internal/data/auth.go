package data

import (
	"context"
	"fmt"
	"github.com/realotz/mstore/internal/biz"
	"time"
)

const PasswdErrorCountKey = "PasswdErrorCount"
const CodeKey = "Code"
const CodeLastTime = "Code:LastTime"

func NewAuthRepo(data *Data)biz.AuthRepo {
	return &authRepo{data:data}
}

type authRepo struct {
	data *Data
}

// 获取验证码
func (u *authRepo) GetSmsCode(ctx context.Context, biz, account string) string {
	key := fmt.Sprintf("%s:%s:%s", CodeKey, biz, account)
	if d, ok := u.data.cache.Get(key); ok {
		return d.(string)
	}
	return ""
}

// 获取最近一次短信发送时间戳
func (u *authRepo) LastSmsCodeTime(ctx context.Context, account string) int64 {
	key := fmt.Sprintf("%s:%s", CodeLastTime, account)
	if d, ok := u.data.cache.Get(key); ok {
		return d.(int64)
	}
	return 0
}

// 删除验证码
func (u *authRepo) DelSmsCode(ctx context.Context, biz, account string) {
	key := fmt.Sprintf("%s:%s:%s", CodeKey, biz, account)
	u.data.cache.Delete(key)
}

// 保存验证码
func (u *authRepo) SaveSmsCode(ctx context.Context, biz, account, code string) error {
	key := fmt.Sprintf("%s:%s:%s", CodeKey, biz, account)
	fmt.Println(key)
	lastKey := fmt.Sprintf("%s:%s", CodeLastTime, account)
	u.data.cache.Set(key, code, time.Minute*30)
	u.data.cache.Set(lastKey, time.Now().Unix(), time.Minute*5)
	return nil
}

// 获取密码错误统计
func (u *authRepo) CountPasswdError(ctx context.Context, account string) int {
	key := fmt.Sprintf("%s:%s", PasswdErrorCountKey, account)
	if d, ok := u.data.cache.Get(key); ok {
		return d.(int)
	}
	return 0
}

// 记录密码错误
func (u *authRepo) AddPasswdError(ctx context.Context, account string) {
	key := fmt.Sprintf("%s:%s", PasswdErrorCountKey, account)
	var n = 0
	if d, ok := u.data.cache.Get(key); ok {
		n = d.(int)
	}
	n += 1
	u.data.cache.Set(key, n, time.Hour)
}

// 清除统计
func (u *authRepo) DelPasswdError(ctx context.Context, account string) {
	key := fmt.Sprintf("%s:%s", PasswdErrorCountKey, account)
	u.data.cache.Delete(key)
}

