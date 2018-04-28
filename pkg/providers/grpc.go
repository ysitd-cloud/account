package providers

import (
	"net"

	grpcService "code.ysitd.cloud/auth/account/pkg/grpc"
	"code.ysitd.cloud/auth/account/pkg/grpc/proxy"
	"code.ysitd.cloud/auth/account/pkg/metrics"
	pb "code.ysitd.cloud/grpc/schema/account"
	"github.com/tonyhhyip/go-di-container"
	"golang.ysitd.cloud/db"
	"google.golang.org/grpc"
)

type grpcServiceProvder struct {
	*container.AbstractServiceProvider
}

func (*grpcServiceProvder) Provides() []string {
	return []string{
		"grpc.server",
		"grpc.listener",
		"grpc.proxy",
	}
}

func (*grpcServiceProvder) Register(app container.Container) {
	app.Singleton("grpc.service", func(app container.Container) interface{} {
		pool := app.Make("db.pool").(db.Opener)
		collector := app.Make("metrics").(metrics.Collector)
		service := &grpcService.AccountService{
			Pool:      pool,
			Container: app,
			Collector: collector,
		}

		service.Init()

		return service
	})

	app.Singleton("grpc.listener", func(app container.Container) interface{} {
		listener, err := net.Listen("tcp", "localhost:50051")
		if err != nil {
			panic(err)
		}

		return listener
	})

	app.Singleton("grpc.server", func(app container.Container) interface{} {
		server := grpc.NewServer()

		service := app.Make("grpc.service").(pb.AccountServer)

		pb.RegisterAccountServer(server, service)

		return server
	})

	app.Singleton("grpc.proxy", func(app container.Container) interface{} {
		service := app.Make("grpc.service").(pb.AccountServer)
		return proxy.CreateProxy(service)
	})
}
