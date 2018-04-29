package http

import (
	"code.ysitd.cloud/auth/account/pkg/metrics"
	"code.ysitd.cloud/gin/utils/interfaces"
	"github.com/tonyhhyip/go-di-container"
)

func NewServiceProvider(app container.Container) container.ServiceProvider {
	sp := &serviceProvider{
		AbstractServiceProvider: container.NewAbstractServiceProvider(true),
	}

	sp.SetContainer(app)

	return sp
}

func newService(collector *metrics.Collector, app container.Container) interfaces.Service {
	return &service{
		Collector: collector,
		app:       app,
	}
}
