package http

import (
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/account/http/handler/oauth"
	"github.com/ysitd-cloud/account/http/middlewares"
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
		middlewares.ContainsJudgeHeader,
		middlewares.JudgeToken("validate", "cloud.ysitd.account.token"),
		oauth.ValidateToken,
	)
}
