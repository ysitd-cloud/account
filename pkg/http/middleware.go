package http

import (
	"code.ysitd.cloud/component/account/pkg/http/middlewares"
	"github.com/gin-gonic/gin"
)

func (s *service) bindMiddleware(app gin.IRouter) {
	app.Use(s.app.Make("session.middleware").(gin.HandlerFunc))
	app.Use(middlewares.Security())
}
