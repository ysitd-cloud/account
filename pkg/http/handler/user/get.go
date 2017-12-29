package user

import (
	"net/http"

	"github.com/RangelReale/osin"
	"github.com/gin-gonic/gin"
	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/account/pkg/model/user"
	"github.com/ysitd-cloud/account/pkg/utils"
)

func getUser(c *gin.Context) {
	kernel := c.MustGet("kernel").(container.Kernel)
	db := kernel.Make("db.pool").(utils.DatabasePool)

	id := c.Param("instance")
	instance, err := user.LoadFromDBWithUsername(db, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, instance)
		c.Abort()
	}
}

func getUserInfo(c *gin.Context) {
	access := c.MustGet("oauth.access").(*osin.AccessData)
	approved := access.UserData.(string)

	kernel := c.MustGet("kernel").(container.Kernel)
	db := kernel.Make("db.pool").(utils.DatabasePool)

	instance, err := user.LoadFromDBWithUsername(db, approved)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, instance)
		c.Abort()
	}
}
