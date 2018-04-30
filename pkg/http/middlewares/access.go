package middlewares

import (
	"net/http"

	"code.ysitd.cloud/auth/account/third_party/forked/github.com/RangelReale/osin"
	"github.com/gin-gonic/gin"
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
