package http

import (
	"code.ysitd.cloud/component/account/pkg/http/handler/user"
	"code.ysitd.cloud/component/account/pkg/http/middlewares"
	"github.com/gin-gonic/gin"
)

func registerApi(group *gin.RouterGroup) {
	v1 := group.Group("v1", middlewares.ContainsAuthHeader, middlewares.BearerToken)
	user.Register(v1)
}
