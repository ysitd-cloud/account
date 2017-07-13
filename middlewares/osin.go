package middlewares

import (
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/ysitd-cloud/account/setup"
)

func Osin() gin.HandlerFunc {
	server := setup.SetupOsinServer()
	return func (c *gin.Context) {
		c.Set("osin.server", server)
		c.Next()
	}
}
