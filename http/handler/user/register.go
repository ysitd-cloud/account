package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/account/http/middlewares"
)

func Register(app *gin.RouterGroup) {
	app.GET("/user/info", getUserInfo)
	app.GET("/users",
		middlewares.JudgeToken("list", "cloud.ysitd.account.user"),
		listUsers,
	)
	group := app.Group("/users")
	bindGroup(group)
}

func bindGroup(group *gin.RouterGroup) {
	group.GET("/:user",
		middlewares.JudgeToken("read", "cloud.ysitd.account.user"),
		middlewares.CheckGetUserAccess,
		getUser,
	)
}
