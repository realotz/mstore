package auth

import (
	"context"
	"encoding/base64"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/realotz/mstore/api/errors"
	"github.com/realotz/mstore/pkg/myctx"
	"net/textproto"
	"strings"
)

type CheckUser interface {
	ParseUserToken(ctx context.Context, token string) (*myctx.JwtCustomClaims, error)
}

type FuncAuth func(context.Context, transport.Transporter) (context.Context, error)

// 排除路由
var noAuthRole = map[string]bool{
	"/api.service.users.v1.AuthService/Captcha":   true,
	"/api.service.users.v1.AuthService/Login":     true,
	"/api.service.users.v1.AuthService/NewPasswd": true,
}

var prefixRule = []string{
}

// auth 授权登录认证中间件
func Server(fs ...FuncAuth) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if info, ok := transport.FromServerContext(ctx); ok {
				for _, prefix := range prefixRule {
					if strings.HasPrefix(info.Operation(), prefix) {
						return handler(ctx, req)
					}
				}
				if _, ok = noAuthRole[info.Operation()]; !ok {
					for _, f := range fs {
						ctx, err = f(ctx, info)
						if err == nil {
							return handler(ctx, req)
						}
					}
					return nil, err
				}
			}
			return handler(ctx, req)
		}
	}
}

// 简单解析cookie内容
func parseCookie(str string) map[string]string {
	res := make(map[string]string)
	line := textproto.TrimString(str)
	var part string
	for len(line) > 0 { // continue since we have rest
		if splitIndex := strings.Index(line, ";"); splitIndex > 0 {
			part, line = line[:splitIndex], line[splitIndex+1:]
		} else {
			part, line = line, ""
		}
		part = textproto.TrimString(part)
		if len(part) == 0 {
			continue
		}
		name, val := part, ""
		if j := strings.Index(part, "="); j >= 0 {
			name, val = name[:j], name[j+1:]
		}
		if name == "" {
			continue
		}
		res[name] = val
	}
	return res
}

func UserAuth(c CheckUser) FuncAuth {
	return func(ctx context.Context, info transport.Transporter) (context.Context, error) {
		var token string
		ck := info.RequestHeader().Get("Cookie")
		if ck != "" {
			cks := parseCookie(ck)
			if tk, ok := cks["TOKEN"]; ok {
				bytes, err := base64.URLEncoding.DecodeString(tk)
				if err == nil {
					token = string(bytes)
				}
			}
		}
		if token == "" {
			token = info.RequestHeader().Get("Authorization")
			token = strings.TrimSpace(token)
		}
		if token == "" {
			return ctx, errors.ErrorNotLogin("未携带token身份")
		}
		claims, err := c.ParseUserToken(ctx, token)
		if err != nil {
			return ctx, errors.ErrorNotLogin("身份出错，请重新登录")
		}
		ctx = myctx.WithUserInfoContext(ctx, claims)
		return ctx, nil
	}
}
