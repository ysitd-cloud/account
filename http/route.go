package http

import (
	"github.com/ysitd-cloud/account/http/handler"
	"github.com/ysitd-cloud/account/http/handler/user"
	"github.com/ysitd-cloud/account/http/middlewares"
	"gopkg.in/gin-gonic/gin.v1"
)

func Register(app *gin.Engine) {
	app.Use(middlewares.DB())
	app.Use(middlewares.Sessions())
	app.Use(middlewares.Osin())
	app.Use(middlewares.Judge())

	app.GET("/authorize",
		handler.HandleAuthorize,
		middlewares.LoginOrRedirect,
		handler.HandleAuthorizeApprove,
	)
	app.POST("/authorize",
		handler.HandleAuthorize,
		middlewares.LoginOrRedirect,
		handler.HandleAuthorizeApprove,
	)

	app.POST("/token", handler.HandleTokenRequest)
	app.GET("/login", handler.LoginForm)
	app.POST("/login", handler.LoginPost)

	user.Register(app)
}
