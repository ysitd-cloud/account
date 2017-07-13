package middlewares

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
)

func RequireLogin() gin.HandlerFunc {
	return func (c *gin.Context) {
		session := getSession(c)
		if session.Exists("username") {
			c.Next()
		}

		c.AbortWithStatus(http.StatusForbidden)
	}
}
