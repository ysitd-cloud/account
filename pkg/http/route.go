package http

import (
	"code.ysitd.cloud/component/account/pkg/http/handler"
	"github.com/gin-gonic/gin"
)

func register(app gin.IRouter) {
	if gin.ReleaseMode != gin.Mode() {
		proxy := app.Group("/assets/")
		proxy.GET("/:assets", handler.AssetsProxy)
	}
}
