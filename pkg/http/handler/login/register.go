package login

import (
	"github.com/gin-gonic/gin"
)

const (
	EndpointLoginForm   = "login_form"
	EndpointLoginSubmit = "login_submit"
)

func Register(app gin.IRoutes) {
	app.GET("/login", basicForm)
	app.POST("/login", basicSubmit)
}
