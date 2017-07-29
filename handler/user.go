package handler

import (
	"database/sql"
	"net/http"

	"github.com/RangelReale/osin"
	"github.com/ysitd-cloud/account/model"
	"gopkg.in/gin-gonic/gin.v1"
)

func CheckGetUserAccess(c *gin.Context) {
	username := c.Param("user")
	access := c.MustGet("oauth.access").(*osin.AccessData)
	approved := access.UserData.(string)
	if username != approved {
		c.AbortWithStatus(http.StatusForbidden)
	} else {
		c.Next()
	}
}

func GetUser(c *gin.Context) {
	userId := c.Param("user")
	db := c.MustGet("db").(*sql.DB)
	user, err := model.LoadUserFromDBWithUsername(db, userId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
		c.Abort()
	}
}
