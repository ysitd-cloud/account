package main

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/ysitd-cloud/account/middlewares"
	"github.com/ysitd-cloud/account/setup"
	"github.com/ysitd-cloud/account/handler"
)

func main() {
	server := setup.SetupOsinServer()

	app := gin.Default()
	app.LoadHTMLGlob("views/*.tmpl")

	app.Use(middlewares.DB())
	app.Use(middlewares.Sessions())

	app.GET("/authorize", middlewares.HandleAuthorize(server), handler.LoginForm, middlewares.HandleAuthorizeApprove(server))
	app.POST("/authorize",
		middlewares.HandleAuthorize(server),
		handler.LoginPost,
		handler.LoginForm,
		middlewares.HandleAuthorizeApprove(server),
	)

	app.POST("/token", handler.HandleTokenRequest(server))

	app.Run(":" + os.Getenv("PORT"))
}
