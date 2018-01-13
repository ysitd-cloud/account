package user

import (
	"net/http"

	"code.ysitd.cloud/common/go/db"
	"code.ysitd.cloud/component/account/pkg/model/user"
	"github.com/gin-gonic/gin"
	"github.com/tonyhhyip/go-di-container"
)

func listUsers(c *gin.Context) {
	kernel := c.MustGet("kernel").(container.Kernel)
	db := kernel.Make("db.pool").(db.Pool)

	users, err := user.ListFromDB(db)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, users)
}
