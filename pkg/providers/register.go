package providers

import (
	"code.ysitd.cloud/component/account/pkg/http"
	"code.ysitd.cloud/component/account/pkg/metrics"
	"github.com/tonyhhyip/go-di-container"
)

func Register(kernel container.Kernel) {
	kernel.Register(func(app container.Container) container.ServiceProvider {
		sp := &etcdServiceProvider{
			AbstractServiceProvider: container.NewAbstractServiceProvider(true),
		}
		sp.SetContainer(app)

		return sp
	})
	kernel.Register(func(app container.Container) container.ServiceProvider {
		sp := databaseServiceProvider{
			AbstractServiceProvider: container.NewAbstractServiceProvider(true),
		}
		sp.SetContainer(app)

		return &sp
	})

	kernel.Register(func(app container.Container) container.ServiceProvider {
		sp := redisServiceProvider{
			AbstractServiceProvider: container.NewAbstractServiceProvider(true),
		}
		sp.SetContainer(app)

		return &sp
	})

	kernel.Register(func(app container.Container) container.ServiceProvider {
		sp := osinServiceProvider{
			AbstractServiceProvider: container.NewAbstractServiceProvider(true),
		}
		sp.SetContainer(app)

		return &sp
	})

	kernel.Register(func(app container.Container) container.ServiceProvider {
		sp := sessionServiceProvider{
			AbstractServiceProvider: container.NewAbstractServiceProvider(true),
		}
		sp.SetContainer(app)

		return &sp
	})

	kernel.Register(func(app container.Container) container.ServiceProvider {
		sp := &grpcServiceProvder{
			AbstractServiceProvider: container.NewAbstractServiceProvider(true),
		}

		sp.SetContainer(app)

		return sp
	})

	kernel.Register(metrics.NewServiceProvider)
	kernel.Register(http.NewServiceProvider)
}
