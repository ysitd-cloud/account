package middlewares

import (
	"net/http"
	"strings"

	"code.ysitd.cloud/auth/account/third_party/forked/github.com/RangelReale/osin"
	"github.com/gin-gonic/gin"
	"github.com/tonyhhyip/go-di-container"
)

func BearerToken(c *gin.Context) {

	if c.MustGet("authorization.type").(string) != "bearer" {
		c.Next()
		return
	}

	kernel := c.MustGet("kernel").(container.Kernel)
	server := kernel.Make("osin.server").(*osin.Server)
	defer server.Storage.Close()

	token := c.MustGet("authorization.value").(string)

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
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	pieces := strings.Split(authHeader, " ")
	if len(pieces) != 2 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	authType := strings.ToLower(pieces[0])

	if authType != "bearer" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Set("authorization.type", authType)
	c.Set("authorization.value", pieces[1])

	c.Next()
}
