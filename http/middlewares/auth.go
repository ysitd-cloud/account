package middlewares

import (
	"net/http"
	"strings"

	"github.com/RangelReale/osin"
	"gopkg.in/gin-gonic/gin.v1"
)

func BearerToken(c *gin.Context) {
	token := c.MustGet("authorization.value").(string)
	server := c.MustGet("osin.server").(*osin.Server)

	if c.MustGet("authorization.type").(string) != "bearer" {
		c.Next()
		return
	}

	if access, err := server.Storage.LoadAccess(token); err != nil {
		c.AbortWithStatus(http.StatusForbidden)
	} else {
		c.Set("oauth.access", access)
		c.Next()
	}
}

func ContainsAuthHeader(c *gin.Context) {
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

	authType := pieces[0]

	if authType != "bearer" && authType != "token" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Set("authorization.type", authType)
	c.Set("authorization.value", pieces[1])

	c.Next()
}
