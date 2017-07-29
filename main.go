package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/ysitd-cloud/account/http"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	app := gin.Default()
	app.LoadHTMLGlob("views/*.tmpl")

	http.Register(app)

	app.Run(":" + os.Getenv("PORT"))
}
