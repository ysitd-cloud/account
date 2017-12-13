package providers

import (
	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/account/pkg/grpc"
)

var Kernel container.Kernel = container.NewKernel()

func init() {
	Kernel.Register(func(app container.Container) container.ServiceProvider {
		sp := databaseServiceProvider{
			AbstractServiceProvider: container.NewAbstractServiceProvider(true),
		}
		sp.SetContainer(app)

		return &sp
	})

	Kernel.Register(func(app container.Container) container.ServiceProvider {
		sp := redisServiceProvider{
			AbstractServiceProvider: container.NewAbstractServiceProvider(true),
		}
		sp.SetContainer(app)

		return &sp
	})

	Kernel.Register(func(app container.Container) container.ServiceProvider {
		sp := osinServiceProvider{
			AbstractServiceProvider: container.NewAbstractServiceProvider(true),
		}
		sp.SetContainer(app)

		return &sp
	})

	Kernel.Register(func(app container.Container) container.ServiceProvider {
		sp := sessionServiceProvider{
			AbstractServiceProvider: container.NewAbstractServiceProvider(true),
		}
		sp.SetContainer(app)

		return &sp
	})

	Kernel.Register(grpc.CreateGrpcServiceProvider)
}
