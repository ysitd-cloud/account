package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/ysitd-cloud/account/handler"
	"github.com/ysitd-cloud/account/middlewares"
	"gopkg.in/gin-gonic/gin.v1"
	"os"
)

func main() {
	app := gin.Default()
	app.LoadHTMLGlob("views/*.tmpl")

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

	app.GET("/user/:user",
		middlewares.AuthToken,
		handler.CheckGetUserAccess,
		handler.GetUser,
	)

	app.Run(":" + os.Getenv("PORT"))
}
