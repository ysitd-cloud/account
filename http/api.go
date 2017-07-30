package http

import (
	"github.com/ysitd-cloud/account/http/handler/user"
	"github.com/ysitd-cloud/account/http/middlewares"
	"gopkg.in/gin-gonic/gin.v1"
)

func registerApi(group *gin.RouterGroup) {
	v1 := group.Group("v1", middlewares.ContainsAuthHeader, middlewares.BearerToken)
	user.Register(v1)
}
