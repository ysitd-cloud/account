package http

import (
	"code.ysitd.cloud/component/account/pkg/http/handler"
	"code.ysitd.cloud/component/account/pkg/http/handler/login"
	"code.ysitd.cloud/component/account/pkg/http/handler/pages"
	"code.ysitd.cloud/component/account/pkg/http/middlewares"
	"code.ysitd.cloud/component/account/pkg/kernel"
	"code.ysitd.cloud/component/account/pkg/metrics"
	"github.com/gin-gonic/gin"
)

func Register(app gin.IRouter, collector metrics.Collector) {
	app.Use(middlewares.BindKernel)
	app.Use(kernel.Kernel.Make("session.middleware").(gin.HandlerFunc))
	app.Use(middlewares.Security())
	login.Register(app, collector)
	pages.Register(app)
	{
		api := app.Group("/api")
		registerApi(api)
	}

	{
		oauth := app.Group("/oauth")
		registerOAuth(oauth)
	}

	if gin.ReleaseMode != gin.Mode() {
		proxy := app.Group("/assets/")
		proxy.GET("/:assets", handler.AssetsProxy)
	}
}
