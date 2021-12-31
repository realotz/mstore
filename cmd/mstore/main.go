package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/realotz/mstore/internal/conf"
	"github.com/realotz/mstore/pkg/cron"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"os"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}


// 链路追踪
func newTrace(cfg *conf.Data) {
	if cfg.Jaeger!=nil{
		exp, err := jaeger.New(jaeger.WithAgentEndpoint(
			jaeger.WithAgentHost(cfg.Jaeger.Host),
			jaeger.WithAgentPort(cfg.Jaeger.Port)))
		if err != nil {
			return
		}
		tp := tracesdk.NewTracerProvider(
			tracesdk.WithBatcher(exp),
			tracesdk.WithResource(resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(Name),
			)),
		)
		otel.SetTracerProvider(tp)
	}
}


func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server, cron *cron.CronManager) *kratos.App {
	return kratos.New(
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
			cron,
		),
	)
}

func main() {
	flag.Parse()
	logger := log.NewStdLogger(os.Stdout)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app,cleanup, err := initApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	newTrace(bc.Data)
	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
