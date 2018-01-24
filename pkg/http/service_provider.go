package http

import (
	"code.ysitd.cloud/component/account/pkg/metrics"
	"github.com/tonyhhyip/go-di-container"
)

type serviceProvider struct {
	*container.AbstractServiceProvider
}

func (*serviceProvider) Provides() []string {
	return []string{
		"http.service",
	}
}

func (*serviceProvider) Register(app container.Container) {
	app.Bind("http.service", func(app container.Container) interface{} {
		collector := app.Make("metrics").(metrics.Collector)
		s := newService(collector, app).(*service)
		s.init()
		return s
	})
}
