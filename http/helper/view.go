package helper

import (
	"fmt"
	"os"

	"github.com/CloudyKit/jet"
	"gopkg.in/gin-gonic/gin.v1"
)

var host string = os.Getenv("STATIC_ADDRESS")

func RenderAppView(c *gin.Context, code int, view, title string) {
	vars := make(jet.VarMap)
	staticPath := fmt.Sprintf("%s/assets", host)
	vars.Set("title", title)
	vars.Set("view", view)
	vars.Set("script", fmt.Sprintf("%s/%s/app.js", staticPath, view))
	vars.Set("style", fmt.Sprintf("%s/%s/app.css", staticPath, view))
	vars.Set("staticPath", staticPath)
	c.HTML(code, "app.jet", vars)
}
