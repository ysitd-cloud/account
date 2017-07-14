package middlewares

import (
	"os"

	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ysitd-cloud/account/setup"
	sessions "github.com/ysitd-cloud/gin-sessions"
)

func Sessions() gin.HandlerFunc {
	store, err := setup.SetupSessionStore()
	if err != nil {
		panic(err)
	}

	name := os.Getenv("SESSION_NAME")

	return sessions.Sessions(name, store, true)
}

func GetSession(c *gin.Context) sessions.Session {
	return sessions.Default(c)
}
