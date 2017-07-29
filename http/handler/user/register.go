package user

import (
	"github.com/ysitd-cloud/account/http/middlewares"
	"gopkg.in/gin-gonic/gin.v1"
)

func Register(app *gin.Engine) {
	group := app.Group("/user", middlewares.ContainsAuthHeader)
	bindGroup(group)
}

func bindGroup(group *gin.RouterGroup) {
	group.GET("/:user",
		middlewares.BearerToken,
		middlewares.CheckGetUserAccess,
		getUser,
	)
}
