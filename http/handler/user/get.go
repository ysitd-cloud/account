package user

import (
	"database/sql"
	"net/http"

	"github.com/ysitd-cloud/account/model"
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/RangelReale/osin"
)

func getUser(c *gin.Context) {
	id := c.Param("user")
	db := c.MustGet("db").(*sql.DB)
	user, err := model.LoadUserFromDBWithUsername(db, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
		c.Abort()
	}
}

func getUserInfo(c *gin.Context) {
	access := c.MustGet("oauth.access").(*osin.AccessData)
	approved := access.UserData.(string)
	db := c.MustGet("db").(*sql.DB)
	user, err := model.LoadUserFromDBWithUsername(db, approved)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, user)
		c.Abort()
	}
}
