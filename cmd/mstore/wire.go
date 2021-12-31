// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/realotz/mstore/internal/conf"
	"github.com/realotz/mstore/internal/data"
	"github.com/realotz/mstore/internal/server"
	"github.com/realotz/mstore/internal/service"
	"github.com/realotz/mstore/internal/biz"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error)  {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
