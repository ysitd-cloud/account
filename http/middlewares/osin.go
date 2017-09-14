package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/account/setup"
)

func Osin() gin.HandlerFunc {
	return func(c *gin.Context) {
		server := setup.SetupOsinServer()
		defer server.Storage.Close()
		c.Set("osin.server", server)
		c.Next()
	}
}
