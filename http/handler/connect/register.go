package connect

import (
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/account/http/middlewares"
)

func Register(app *gin.Engine) {
	group := app.Group("/connect", middlewares.LoginOrRedirect)
	group.GET("/:provider", middlewares.LoginOrRedirect, redirectToOAuth)
	group.GET("/:provider/callback", middlewares.LoginOrRedirect, oauthCallback)
}
