package user

import (
	"database/sql"
	"net/http"

	"github.com/ysitd-cloud/account/model"
	"gopkg.in/gin-gonic/gin.v1"
)

func listUsers(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	users, err := model.ListUserFromDB(db)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, users)
}
