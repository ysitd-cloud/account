package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/account/setup"
)

func Judge() gin.HandlerFunc {

	client := setup.SetupJudgeClient()

	return func(c *gin.Context) {
		c.Set("judge", client)
		c.Next()
	}
}
