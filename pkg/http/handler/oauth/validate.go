package oauth

import (
	"net/http"

	"github.com/RangelReale/osin"
	"github.com/gin-gonic/gin"
	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/account/pkg/model/user"
	"github.com/ysitd-cloud/account/pkg/utils"
)

func ValidateToken(c *gin.Context) {
	kernel := c.MustGet("kernel").(container.Kernel)
	server := kernel.Make("osin.server").(*osin.Server)
	defer server.Storage.Close()

	token := c.Query("token")
	if token == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	access, err := server.Storage.LoadAccess(token)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	id := access.UserData.(string)
	db := kernel.Make("db.pool").(utils.DatabasePool)
	instance, err := user.LoadFromDBWithUsername(db, id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.AbortWithStatusJSON(http.StatusOK, instance)
	}
}
