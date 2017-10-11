package http

import (
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/account/http/handler"
	"github.com/ysitd-cloud/account/http/handler/connect"
	"github.com/ysitd-cloud/account/http/handler/login"
	"github.com/ysitd-cloud/account/http/handler/pages"
	"github.com/ysitd-cloud/account/http/middlewares"
	"github.com/ysitd-cloud/account/providers"
)

func Register(app *gin.Engine) {
	app.Use(middlewares.BindKernel)
	app.Use(providers.Kernel.Make("session.middleware").(gin.HandlerFunc))
	app.Use(middlewares.Security())
	login.Register(app)
	connect.Register(app)
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
