package providers

import (
	"github.com/sirupsen/logrus"
	"github.com/tonyhhyip/go-di-container"
)

type loggerServiceProvider struct {
	*container.AbstractServiceProvider
}

func (*loggerServiceProvider) Provides() []string {
	return []string{
		"logger",
	}
}

func (*loggerServiceProvider) Register(app container.Container) {
	app.Singleton("logger", func(app container.Container) interface{} {
		return logrus.New()
	})
}
