package user

import (
	"net/http"

	"code.ysitd.cloud/auth/account/pkg/model/user"
	"github.com/gin-gonic/gin"
	"github.com/tonyhhyip/go-di-container"
	"golang.ysitd.cloud/db"
)

func listUsers(c *gin.Context) {
	kernel := c.MustGet("kernel").(container.Kernel)
	db := kernel.Make("db.pool").(db.Opener)

	users, err := user.ListFromDB(db)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, users)
}
