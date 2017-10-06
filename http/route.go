package http

import (
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/account/http/handler"
	"github.com/ysitd-cloud/account/http/middlewares"
	"github.com/ysitd-cloud/account/providers"
)

func Register(app *gin.Engine) {
	app.Use(middlewares.BindKernel)
	app.Use(providers.Kernel.Make("session.middleware").(gin.HandlerFunc))
	app.GET("/login", handler.LoginForm)
	app.POST("/login", handler.LoginPost)

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
