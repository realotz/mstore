package rbac

import (
	"context"
	"git.hxecloud.com/cloudwonder-portal/rms-server/pkg/myctx"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/realotz/mstore/api/errors"
	"strings"
)

// admin 权限拦截
func Server() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			userInfo, err := myctx.FormUserInfo(ctx)
			if err != nil {
				return nil, errors.ErrorNotLogin("无非获取登录信息")
			}
			tr, ok := transport.FromServerContext(ctx)
			if ok {
				if strings.Contains(tr.Operation(), "/api.agent.v1.AgentService") {
					if userInfo.Role != 9 {
						return nil, errors.ErrorNotAuthority("无权限的角色")
					}
				} else {
					if userInfo.Role != 1 && userInfo.Role != 2 {
						return nil, errors.ErrorNotAuthority("无权限的角色")
					}
				}
			}
			return handler(ctx, req)
		}
	}
}
