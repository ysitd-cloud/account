package http

import (
	"code.ysitd.cloud/component/account/pkg/metrics"
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

func newService(collector metrics.Collector) interfaces.Service {
	return &service{
		collector: collector,
	}
}
