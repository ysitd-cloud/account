package middlewares

import (
	"os"

	"gopkg.in/gin-gonic/gin.v1"

	"github.com/ysitd-cloud/account/setup"
	sessions "github.com/ysitd-cloud/gin-sessions"
)

var SESSION_NAME string = os.Getenv("SESSION_NAME")

func Sessions() gin.HandlerFunc {
	store, err := setup.SetupSessionStore()
	if err != nil {
		panic(err)
	}

	return sessions.Sessions(SESSION_NAME, store, true)
}

func GetSession(c *gin.Context) sessions.Session {
	return sessions.Default(c)
}
