package middlewares

import (
	"github.com/ysitd-cloud/account/setup"
	"gopkg.in/gin-gonic/gin.v1"
)

func Judge() gin.HandlerFunc {

	client := setup.SetupJudgeClient()

	return func(c *gin.Context) {
		c.Set("judge", client)
		c.Next()
	}
}
