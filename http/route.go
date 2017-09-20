package http

import (
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/account/http/handler/connect"
	"github.com/ysitd-cloud/account/http/handler/login"
	"github.com/ysitd-cloud/account/http/middlewares"
)

func Register(app *gin.Engine) {
	app.Use(middlewares.DB())
	app.Use(middlewares.Sessions())
	app.Use(middlewares.Osin())
	app.Use(middlewares.Judge())
	login.Register(app)
	connect.Register(app)
	{
		api := app.Group("/api")
		registerApi(api)
	}

	{
		oauth := app.Group("/oauth")
		registerOAuth(oauth)
	}
}
