package main

import (
	"os"

	"github.com/CloudyKit/jet"
	_ "github.com/joho/godotenv/autoload"
	"github.com/ysitd-cloud/account/http"
	render "github.com/ysitd-cloud/gin-jet"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	app := gin.Default()
	views := jet.NewHTMLSet("./views")
	template := render.NewJetRender(views)
	app.HTMLRender = template

	http.Register(app)

	app.Run(":" + os.Getenv("PORT"))
}
