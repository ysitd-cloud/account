package http

import (
	"github.com/ysitd-cloud/account/http/handler"
	"github.com/ysitd-cloud/account/http/middlewares"
	"gopkg.in/gin-gonic/gin.v1"
)

func userGroupRoute(group *gin.RouterGroup) {
	group.GET("/:user",
		middlewares.BearerToken,
		middlewares.CheckGetUserAccess,
		handler.GetUser,
	)
}

func Register(app *gin.Engine) {
	app.Use(middlewares.DB())
	app.Use(middlewares.Sessions())
	app.Use(middlewares.Osin())

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

	{
		userRoute := app.Group("/user", middlewares.ContainsAuthHeader)
		userGroupRoute(userRoute)
	}
}
