package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/account/pkg/kernel"
)

func BindKernel(c *gin.Context) {
	c.Set("kernel", kernel.Kernel)
	c.Next()
}
