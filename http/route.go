package http

import (
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/account/http/handler"
	"github.com/ysitd-cloud/account/http/middlewares"
)

func Register(app *gin.Engine) {
	app.Use(middlewares.Security())
	app.Use(middlewares.DB())
	app.Use(middlewares.Sessions())
	app.Use(middlewares.Osin())
	app.Use(middlewares.Judge())
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
