package metrics

import (
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
	app.Singleton("metrics", func(app container.Container) interface{} {
		return NewCollector()
	})
}
