package http

import (
	"github.com/ysitd-cloud/account/http/handler/user"
	"gopkg.in/gin-gonic/gin.v1"
)

func registerApi(group *gin.RouterGroup) {
	v1 := group.Group("v1")
	user.Register(v1)
}
