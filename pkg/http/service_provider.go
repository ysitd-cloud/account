package http

import (
	"code.ysitd.cloud/auth/account/pkg/http/handler"
	"code.ysitd.cloud/auth/account/pkg/metrics"
	"github.com/tonyhhyip/go-di-container"
	"golang.ysitd.cloud/db"
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
		opener := app.Make("db").(*db.GeneralOpener)
		s := newService(collector, app).(*service)
		s.init()
		s.LoginHandler = &handler.LoginHandler{
			Opener:    opener,
			Collector: collector,
		}
		return s
	})
}
