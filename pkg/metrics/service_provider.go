package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tonyhhyip/go-di-container"
)

type serviceProvider struct {
	*container.AbstractServiceProvider
}

func (*serviceProvider) Provides() []string {
	return []string{
		"metrics",
	}
}

func (*serviceProvider) Register(app container.Container) {
	app.Instance("metrics.registerer", prometheus.DefaultRegisterer)
	app.Singleton("metrics", func(app container.Container) interface{} {
		registry := app.Make("metrics.registerer").(*prometheus.Registry)
		return NewCollector(registry)
	})
}
