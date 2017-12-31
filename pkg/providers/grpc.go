package providers

import (
	"net"

	"github.com/tonyhhyip/go-di-container"
	grpcService "github.com/ysitd-cloud/account/pkg/grpc"
	"github.com/ysitd-cloud/account/pkg/grpc/proxy"
	"github.com/ysitd-cloud/account/pkg/utils"
	pb "github.com/ysitd-cloud/grpc-schema/account"
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
		pool := app.Make("db.pool").(utils.DatabasePool)
		return &grpcService.AccountService{
			Pool:      pool,
			Container: app,
		}
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
