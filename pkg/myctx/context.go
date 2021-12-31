package myctx

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-kratos/kratos/v2/errors"
	"net/http"
)

type ctxUserInfoKey struct{}
type ctxHttpKey struct{}

type JwtProfile struct {
	ID   uint32 `json:"id"`
	Role string `json:"r"`
}

type JwtCustomClaims struct {
	JwtProfile
	UserType string /*client type*/
	jwt.StandardClaims
}

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func NewHttpServerContext(ctx context.Context, my *Context) context.Context {
	return context.WithValue(ctx, ctxHttpKey{}, my)
}

// 获取登录信息
func FormHttpContext(ctx context.Context) (*Context, error) {
	tr, ok := ctx.Value(ctxHttpKey{}).(*Context)
	if !ok {
		return nil, fmt.Errorf("is not http")
	}
	return tr, nil
}

func WithUserInfoContext(ctx context.Context, claims *JwtCustomClaims) context.Context {
	return context.WithValue(ctx, ctxUserInfoKey{}, claims)
}

// 获取登录信息
func FormUserInfo(ctx context.Context) (*JwtCustomClaims, error) {
	tr, ok := ctx.Value(ctxUserInfoKey{}).(*JwtCustomClaims)
	if !ok {
		return nil, errors.Unauthorized("NOT_LOGIN", "未登录")
	}
	return tr, nil
}

// 获取用户id
func FormUserId(ctx context.Context) (uint32, error) {
	user, err := FormUserInfo(ctx)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}
