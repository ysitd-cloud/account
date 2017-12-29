package http

import (
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/account/pkg/http/handler/user"
	"github.com/ysitd-cloud/account/pkg/http/middlewares"
)

func registerApi(group *gin.RouterGroup) {
	v1 := group.Group("v1", middlewares.ContainsAuthHeader, middlewares.BearerToken)
	user.Register(v1)
}
