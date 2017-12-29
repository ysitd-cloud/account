package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/account/pkg/model/user"
	"github.com/ysitd-cloud/account/pkg/utils"
)

func listUsers(c *gin.Context) {
	kernel := c.MustGet("kernel").(container.Kernel)
	db := kernel.Make("db.pool").(utils.DatabasePool)

	users, err := user.ListFromDB(db)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, users)
}
