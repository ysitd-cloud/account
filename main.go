package main

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/ysitd-cloud/account/middlewares"
	"github.com/ysitd-cloud/account/setup"
	"github.com/ysitd-cloud/account/routes"
)

func main() {
	server := setup.SetupOsinServer()

	app := gin.Default()
	app.LoadHTMLGlob("views/*.tmpl")

	app.GET("/authorize", middlewares.HandleAuthorize(server), routes.LoginForm)

	app.Run(":" + os.Getenv("PORT"))
}
