package middlewares

import (
	"github.com/gin-gonic/gin"
	"golang.ysitd.cloud/gin/sessions"
)

func GetSession(c *gin.Context) sessions.Session {
	return sessions.Default(c)
}
