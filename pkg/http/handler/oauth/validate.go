package oauth

import (
	"net/http"

	"code.ysitd.cloud/auth/account/pkg/model/user"
	"github.com/RangelReale/osin"
	"github.com/gin-gonic/gin"
	"github.com/tonyhhyip/go-di-container"
	"golang.ysitd.cloud/db"
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
	opener := kernel.Make("db.pool").(db.Opener)
	instance, err := user.LoadFromDBWithUsername(c.Request.Context(), opener, id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.AbortWithStatusJSON(http.StatusOK, instance)
	}
}
