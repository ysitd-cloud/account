package user

import (
	"net/http"

	"code.ysitd.cloud/component/account/pkg/model/user"
	"code.ysitd.cloud/component/account/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/tonyhhyip/go-di-container"
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
