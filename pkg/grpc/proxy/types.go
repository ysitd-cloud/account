package proxy

import (
	"code.ysitd.cloud/gin/utils/interfaces"
	"code.ysitd.cloud/grpc/schema/account"
	"github.com/gin-gonic/gin"
)

func CreateProxy(service account.AccountServer) interfaces.Service {
	return &proxy{
		service: service,
	}
}

type proxy struct {
	service account.AccountServer
}

func (p *proxy) createMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("service", p.service)
	}
}

func (p *proxy) CreateService() (app interfaces.Engine) {
	app = gin.Default()
	app.Use(p.createMiddleware())
	app.GET("/token/:token", getTokenInfo)
	app.GET("/user/:username", getUserInfo)
	app.POST("/validate", validateUserPassword)
	return app
}
