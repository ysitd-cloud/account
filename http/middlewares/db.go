package middlewares

import (
	"github.com/ysitd-cloud/account/setup"
	"gopkg.in/gin-gonic/gin.v1"
)

func DB() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := setup.SetupDB()
		if err != nil {
			panic(err)
		}

		defer db.Close()
		c.Set("db", db)
		c.Next()
	}
}
