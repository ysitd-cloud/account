package provider

import (
	"database/sql"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/account/model"
)

func listProvider(c *gin.Context) {
	db := c.MustGet("db").(*sql.DB)
	providers, err := model.ListProvider(db)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, providers)
	}
}
