package main

import (
	"os"

	"github.com/CloudyKit/jet"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/ysitd-cloud/account/http"
	render "github.com/ysitd-cloud/gin-jet"
)

func main() {
	app := gin.Default()

	if gin.Mode() != gin.ReleaseMode {
		pprof.Register(app, nil)
	}

	views := jet.NewHTMLSet("./views")
	template := render.NewJetRender(views)
	app.HTMLRender = template

	http.Register(app)

	app.Run(":" + os.Getenv("PORT"))
}
