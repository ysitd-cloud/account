package user

import (
	"database/sql"
	"net/http"

	"github.com/ysitd-cloud/account/model"
	"gopkg.in/gin-gonic/gin.v1"
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
