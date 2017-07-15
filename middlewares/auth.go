package middlewares

import (
	"net/url"
	"net/http"
	"strings"
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/RangelReale/osin"
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

func AuthToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	pieces := strings.Split(authHeader, " ")
	if len(pieces) != 2 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	token := pieces[1]
	server := c.MustGet("osin.server").(*osin.Server)
	if _, err := server.Storage.LoadAccess(token); err != nil {
		c.AbortWithStatus(http.StatusForbidden)
	} else {
		c.Next()
	}
}
