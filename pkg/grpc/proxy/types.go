package proxy

import (
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/grpc-schema/account"
)

type GrpcProxy interface {
	CreateApp() *gin.Engine
}

func CreateProxy(service account.AccountServer) GrpcProxy {
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

func (p *proxy) CreateApp() (app *gin.Engine) {
	app = gin.Default()
	app.Use(p.createMiddleware())
	app.GET("/token/:token", getTokenInfo)
	app.GET("/user/:username", getUserInfo)
	app.POST("/validate", validateUserPassword)
	return app
}
