package oauth

import (
	"database/sql"
	"net/http"

	"time"

	"github.com/RangelReale/osin"
	"github.com/gin-gonic/gin"
	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/account/model"
)

type tokenInfo struct {
	User      *model.User `json:"user"`
	Scope     string      `json:"scope"`
	ExpireIn  int32       `json:"expire_in"`
	CreatedAt time.Time   `json:"created_at"`
}

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
	db := kernel.Make("db.postgres").(*sql.DB)
	user, err := model.LoadUserFromDBWithUsername(db, id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.AbortWithStatusJSON(http.StatusOK, user)
	}
}
