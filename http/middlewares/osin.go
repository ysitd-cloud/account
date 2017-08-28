package middlewares

import (
	"github.com/ysitd-cloud/account/setup"
	"gopkg.in/gin-gonic/gin.v1"
)

func Osin() gin.HandlerFunc {
	return func(c *gin.Context) {
		server := setup.SetupOsinServer()
		c.Set("osin.server", server)
		c.Next()
	}
}
