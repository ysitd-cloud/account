package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/account/providers"
)

func BindKernel(c *gin.Context) {
	c.Set("kernel", providers.Kernel)
	c.Next()
}
