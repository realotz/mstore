package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	storageV1 "github.com/realotz/mstore/api/storage/v1"
	userV1 "github.com/realotz/mstore/api/users/v1"
	"github.com/realotz/mstore/internal/biz"
	"github.com/realotz/mstore/internal/conf"
	"github.com/realotz/mstore/internal/service/storage"
	"github.com/realotz/mstore/pkg/middleware/auth"
	"github.com/realotz/mstore/pkg/middleware/logging"
	"github.com/realotz/mstore/pkg/middleware/rbac"
	"github.com/realotz/mstore/pkg/middleware/tracing"
	"github.com/realotz/mstore/pkg/middleware/transaction"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, logger log.Logger,
	authUc *biz.AuthUseCase,
	authV1 userV1.AuthServiceServer,
	user userV1.UserServiceServer,
	volume *storage.VolumeService,
) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			metrics.Server(),
			validate.Validator(),
			// 用户认证
			auth.Server(
				auth.UserAuth(authUc),
			),
			// 权限
			selector.Server(rbac.Server()).
				Prefix("/api.admin.v1.").Build(),
			transaction.Server(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	userV1.RegisterUserServiceServer(srv, user)
	userV1.RegisterAuthServiceServer(srv, authV1)
	storageV1.RegisterVolumeServiceServer(srv, volume)
	return srv
}
