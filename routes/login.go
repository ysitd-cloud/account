package routes

import (
	"fmt"
	"net/http"
	"gopkg.in/gin-gonic/gin.v1"
	"html/template"
)

func LoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"url": template.URL(fmt.Sprintf("/authorize?%s", c.Request.URL.RawQuery)),
	})
}
