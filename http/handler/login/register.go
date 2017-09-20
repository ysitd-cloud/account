package login

import (
	"github.com/gin-gonic/gin"
)

func Register(app *gin.Engine) {
	app.GET("/login", basicForm)
	app.POST("/login", basicSubmit)
	group := app.Group("/login")
	group.GET("/:provider", redirectToOAuth)
	group.POST("/:provider/callback", oauthLoginCallback)
}
