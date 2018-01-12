package user

import (
	"code.ysitd.cloud/component/account/pkg/http/middlewares"
	"github.com/gin-gonic/gin"
)

func Register(app *gin.RouterGroup) {
	app.GET("/user/info", getUserInfo)
	app.GET("/users",
		listUsers,
	)
	group := app.Group("/users")
	bindGroup(group)
}

func bindGroup(group *gin.RouterGroup) {
	group.GET("/:user",
		middlewares.CheckGetUserAccess,
		getUser,
	)
}
