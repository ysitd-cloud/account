package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/account/pkg/http/middlewares"
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
