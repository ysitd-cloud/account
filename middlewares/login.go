package middlewares

import (
	"net/http"
	"gopkg.in/gin-gonic/gin.v1"
)

func RequireLogin() gin.HandlerFunc {
	return func (c *gin.Context) {
		session := GetSession(c)
		if session.Exists("username") {
			c.Next()
		}

		c.AbortWithStatus(http.StatusForbidden)
	}
}
