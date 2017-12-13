package pages

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/account/pkg/http/helper"
	"github.com/ysitd-cloud/account/pkg/http/middlewares"
	"github.com/ysitd-cloud/account/pkg/model"
)

func modifiedPassword(c *gin.Context) {
	helper.RenderAppView(c, http.StatusOK, "password", nil)
}

func updatePassword(c *gin.Context) {
	defer c.Abort()

	session := middlewares.GetSession(c)
	username := session.Get("username").(string)

	kernel := c.MustGet("kernel").(container.Kernel)
	db := kernel.Make("db").(*sql.DB)
	defer db.Close()

	user, err := model.LoadUserFromDBWithUsername(db, username)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	origin := c.PostForm("origin")

	if !user.ValidatePassword(origin) {
		c.Redirect(http.StatusFound, "/password?error=password")
		return
	}

	newPassword := c.PostForm("newPassword")

	if newPassword == origin {
		c.Redirect(http.StatusFound, "/password?error=same")
		return
	}

	confirmPassword := c.PostForm("confirmPassword")

	if newPassword != confirmPassword {
		c.Redirect(http.StatusFound, "/password?error=confirm")
		return
	}

	if err = user.ChangePassword(newPassword); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Redirect(http.StatusFound, "/?message=password_change")
}
