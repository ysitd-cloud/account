package pages

import (
	"code.ysitd.cloud/component/account/pkg/http/middlewares"
	"github.com/gin-gonic/gin"
)

func Register(app gin.IRoutes) {
	app.GET("/", middlewares.LoginOrRedirect, profile)
	app.GET("/password", middlewares.LoginOrRedirect, modifiedPassword)
	app.POST("/password", middlewares.LoginOrRedirect, updatePassword)
}
