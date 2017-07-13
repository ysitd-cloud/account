package main

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/ysitd-cloud/account/middlewares"
	"github.com/ysitd-cloud/account/handler"
)

func main() {
	app := gin.Default()
	app.LoadHTMLGlob("views/*.tmpl")

	app.Use(middlewares.DB())
	app.Use(middlewares.Sessions())
	app.Use(middlewares.Osin())

	app.GET("/authorize",
		handler.HandleAuthorize,
		handler.LoginForm,
		handler.HandleAuthorizeApprove,
	)
	app.POST("/authorize",
		handler.HandleAuthorize,
		handler.LoginPost,
		handler.LoginForm,
		handler.HandleAuthorizeApprove,
	)

	app.POST("/token", handler.HandleTokenRequest)

	app.Run(":" + os.Getenv("PORT"))
}
