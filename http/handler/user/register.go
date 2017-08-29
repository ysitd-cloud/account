package user

import (
	"github.com/ysitd-cloud/account/http/middlewares"
	"gopkg.in/gin-gonic/gin.v1"
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
