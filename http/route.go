package http

import (
	"github.com/ysitd-cloud/account/http/handler"
	"github.com/ysitd-cloud/account/http/middlewares"
	"gopkg.in/gin-gonic/gin.v1"
)

func Register(app *gin.Engine) {
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
		registerApi(oauth)
	}
}
