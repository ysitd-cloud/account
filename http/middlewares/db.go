package middlewares

import (
	"github.com/ysitd-cloud/account/setup"
	"gopkg.in/gin-gonic/gin.v1"
)

func DB() gin.HandlerFunc {
	db, err := setup.SetupDB()
	if err != nil {
		panic(err)
	}

	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
