package pages

import (
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/account/http/middlewares"
)

func Register(app *gin.Engine) {
	app.GET("/", middlewares.LoginOrRedirect, profile)
	app.GET("/password", middlewares.LoginOrRedirect, modifiedPassword)
}
