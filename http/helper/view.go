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
	vars.Set("title", title)
	vars.Set("view", view)
	vars.Set("script", fmt.Sprintf("%s/%s/app.js", host, view))
	vars.Set("style", fmt.Sprintf("%s/%s/app.css", host, view))
	c.HTML(code, "app.jet", vars)
}
