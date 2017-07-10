package main

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/ysitd-cloud/account/middlewares"
	"github.com/ysitd-cloud/account/setup"
)

func main() {
	server := setup.SetupOsinServer()

	app := gin.Default()

	app.GET("/authorize", middlewares.HandleAuthorize(server))

	app.Run(":" + os.Getenv("PORT"))
}
