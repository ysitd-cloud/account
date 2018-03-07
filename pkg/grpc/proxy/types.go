package proxy

import (
	"code.ysitd.cloud/grpc/schema/account"
	"github.com/gin-gonic/gin"
	"github.com/tonyhhyip/vodka"
)

func CreateProxy(service account.AccountServer) *proxy {
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

func (p *proxy) CreateService() *vodka.Router {
	app := vodka.NewRouter()
	app.GET("/token/:token", p.getTokenInfo)
	app.GET("/user/:username", p.getUserInfo)
	app.POST("/validate", p.validateUserPassword)
	return app
}
