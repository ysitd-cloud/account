package http

import (
	"code.ysitd.cloud/component/account/pkg/http/handler/oauth"
	"code.ysitd.cloud/component/account/pkg/http/middlewares"
	"github.com/gin-gonic/gin"
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

	group.GET("/token/validate",
		oauth.ValidateToken,
	)

	group.GET("/provider/:provider/callback",
		middlewares.LoginOrRedirect,
		oauth.Callback,
	)
}
