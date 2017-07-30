package user

import (
	"github.com/ysitd-cloud/account/http/middlewares"
	"gopkg.in/gin-gonic/gin.v1"
)

func Register(app *gin.Engine) {
	app.GET("/user",
		middlewares.ContainsAuthHeader,
		middlewares.BearerToken,
		middlewares.JudgeToken("list", "cloud.ysitd.account.user"),
		listUsers,
	)
	group := app.Group("/user", middlewares.ContainsAuthHeader)
	bindGroup(group)
}

func bindGroup(group *gin.RouterGroup) {
	group.GET("/:user",
		middlewares.BearerToken,
		middlewares.JudgeToken("read", "cloud.ysitd.account.user"),
		middlewares.CheckGetUserAccess,
		getUser,
	)
}
