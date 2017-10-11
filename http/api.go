package http

import (
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/account/http/handler/user"
	"github.com/ysitd-cloud/account/http/middlewares"
	"github.com/ysitd-cloud/account/http/handler/provider"
)

func registerApi(group *gin.RouterGroup) {
	v1 := group.Group("v1", middlewares.ContainsAuthHeader, middlewares.BearerToken)
	user.Register(v1)
	provider.Register(group.Group("v1"))
}
