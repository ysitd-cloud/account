package http

import (
	"github.com/ysitd-cloud/account/http/handler/oauth"
	"github.com/ysitd-cloud/account/http/middlewares"
	"gopkg.in/gin-gonic/gin.v1"
)

func registerOAuth(group *gin.RouterGroup) {
	group.GET("/authorize",
		oauth.HandleAuthorize,
		middlewares.LoginOrRedirect,
		oauth.HandleAuthorizeApprove,
	)
	group.POST("/authorize",
		oauth.HandleAuthorize,
		middlewares.LoginOrRedirect,
		oauth.HandleAuthorizeApprove,
	)

	group.POST("/token", oauth.HandleTokenRequest)
}
