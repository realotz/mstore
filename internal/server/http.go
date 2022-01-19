package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	storageV1 "github.com/realotz/mstore/api/storage/v1"
	userV1 "github.com/realotz/mstore/api/users/v1"
	"github.com/realotz/mstore/internal/biz"
	"github.com/realotz/mstore/internal/conf"
	"github.com/realotz/mstore/internal/service/storage"
	_ "github.com/realotz/mstore/pkg/encoding/form"
	"github.com/realotz/mstore/pkg/middleware/auth"
	"github.com/realotz/mstore/pkg/middleware/filter"
	"github.com/realotz/mstore/pkg/middleware/logging"
	"github.com/realotz/mstore/pkg/middleware/tracing"
	"github.com/realotz/mstore/pkg/middleware/transaction"
)

// NewGRPCServer new a gRPC server.
func NewHTTPServer(c *conf.Server, logger log.Logger,
	authUc *biz.AuthUseCase,
	authV1 userV1.AuthServiceServer,
	user userV1.UserServiceServer,
	volume *storage.VolumeService,
) *http.Server {
	var opts = []http.ServerOption{
		http.Filter(filter.HttpContext),
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			metrics.Server(),
			validate.Validator(),
			// 用户认证
			auth.Server(
				auth.UserAuth(authUc),
			),
			//// 权限
			//selector.Server(rbac.Server()).Prefix("/api.admin.v1.", "/api.agent.v1.").Build(),
			transaction.Server(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	userV1.RegisterUserServiceHTTPServer(srv, user)
	userV1.RegisterAuthServiceHTTPServer(srv, authV1)
	storageV1.RegisterVolumeServiceHTTPServer(srv, volume)
	srv.HandlePrefix("/api/v1/big-upload", volume)
	return srv
}
