package middlewares

import (
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/ysitd-cloud/account/setup"
)

func DB() gin.HandlerFunc {
	db, err := setup.SetupDB()
	if err != nil {
		panic(err)
	}

	return func (c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
