package middlewares

import (
	"code.ysitd.cloud/component/account/pkg/kernel"
	"github.com/gin-gonic/gin"
)

func BindKernel(c *gin.Context) {
	c.Set("kernel", kernel.Kernel)
	c.Next()
}
