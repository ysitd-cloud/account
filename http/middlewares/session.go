package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/gin-sessions"
)

func GetSession(c *gin.Context) sessions.Session {
	return sessions.Default(c)
}
