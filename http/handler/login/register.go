package login

import (
	"github.com/gin-gonic/gin"
)

func Register(app *gin.Engine) {
	app.GET("/login", basicForm)
	app.POST("/login", basicSubmit)
}
