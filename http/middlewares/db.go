package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/account/setup"
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
