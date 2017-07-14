package middlewares

import (
	"net/url"
	"net/http"
	"gopkg.in/gin-gonic/gin.v1"
)

func LoginOrRedirect(c *gin.Context) {
	session := GetSession(c)
	if session.Exists("username") {
		c.Next()
		return
	}

	redirect, _ := url.Parse("/login")
	query := redirect.Query()
	query.Set("next", c.Request.URL.String())
	redirect.RawQuery = query.Encode()
	c.Redirect(http.StatusFound, redirect.String())
	c.Abort()
}
