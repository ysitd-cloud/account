package middlewares

import (
	"net/http"

	"github.com/RangelReale/osin"
	"gopkg.in/gin-gonic/gin.v1"
)

func CheckGetUserAccess(c *gin.Context) {
	username := c.Param("user")
	access := c.MustGet("oauth.access").(*osin.AccessData)
	approved := access.UserData.(string)
	if username != approved {
		c.AbortWithStatus(http.StatusForbidden)
	} else {
		c.Next()
	}
}
