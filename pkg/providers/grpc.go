package providers

import (
	"net"
	"net/http"

	grpcService "code.ysitd.cloud/auth/account/pkg/grpc"
	"code.ysitd.cloud/auth/account/pkg/grpc/proxy"
	"code.ysitd.cloud/auth/account/pkg/metrics"
	"code.ysitd.cloud/auth/account/third_party/forked/github.com/RangelReale/osin"
	pb "code.ysitd.cloud/grpc/schema/account"
	"github.com/facebookgo/inject"
	"github.com/tonyhhyip/go-di-container"
	"golang.ysitd.cloud/db"
	"google.golang.org/grpc"
)

var service grpcService.AccountService
var p proxy.Handler

func GetProxyHandler() http.Handler {
	return &p
}

func GetGrpcServer() *grpc.Server {
	server := grpc.NewServer()
	pb.RegisterAccountServer(server, &service)
	return server
}

func InjectGrpcService(graph *inject.Graph) {
	graph.Provide(
		NewObject(&service),
		NewObject(&p),
	)
}

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
		pool := app.Make("db.pool").(*db.GeneralOpener)
		collector := app.Make("metrics").(*metrics.Collector)
		server := app.Make("osin.server").(*osin.Server)
		service := &grpcService.AccountService{
			Pool:      pool,
			Server:    server,
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
		service := app.Make("grpc.service").(*grpcService.AccountService)
		return proxy.CreateProxy(service)
	})
}
