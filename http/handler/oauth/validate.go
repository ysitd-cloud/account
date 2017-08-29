package oauth

import (
	"database/sql"
	"net/http"

	"github.com/RangelReale/osin"
	"github.com/ysitd-cloud/account/model"
	"gopkg.in/gin-gonic/gin.v1"
	"time"
)

type tokenInfo struct {
	User      *model.User `json:"user"`
	Scope     string      `json:"scope"`
	ExpireIn  int32       `json:"expire_in"`
	CreatedAt time.Time   `json:"created_at"`
}

func ValidateToken(c *gin.Context) {
	server := c.MustGet("osin.server").(*osin.Server)
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
	db := c.MustGet("db").(*sql.DB)
	user, err := model.LoadUserFromDBWithUsername(db, id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.AbortWithStatusJSON(http.StatusOK, user)
	}
}
