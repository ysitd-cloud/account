package middlewares

import (
	"github.com/ysitd-cloud/account/setup"
	"gopkg.in/gin-gonic/gin.v1"
)

func Osin() gin.HandlerFunc {
	server := setup.SetupOsinServer()
	return func(c *gin.Context) {
		c.Set("osin.server", server)
		c.Next()
	}
}
