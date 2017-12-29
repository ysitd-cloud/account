package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/account/pkg/http/helper"
	"github.com/ysitd-cloud/account/pkg/http/middlewares"
	"github.com/ysitd-cloud/account/pkg/model/user"
	"github.com/ysitd-cloud/account/pkg/utils"
)

func modifiedPassword(c *gin.Context) {
	helper.RenderAppView(c, http.StatusOK, "password", nil)
}

func updatePassword(c *gin.Context) {
	defer c.Abort()

	session := middlewares.GetSession(c)
	username := session.Get("username").(string)

	kernel := c.MustGet("kernel").(container.Kernel)
	db := kernel.Make("db.pool").(utils.DatabasePool)

	instance, err := user.LoadFromDBWithUsername(db, username)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	origin := c.PostForm("origin")

	if !instance.ValidatePassword(origin) {
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

	if err = instance.ChangePassword(newPassword); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.Redirect(http.StatusFound, "/?message=password_change")
}
