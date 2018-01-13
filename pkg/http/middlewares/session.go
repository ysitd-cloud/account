package middlewares

import (
	"code.ysitd.cloud/gin/sessions"
	"github.com/gin-gonic/gin"
)

func GetSession(c *gin.Context) sessions.Session {
	return sessions.Default(c)
}
