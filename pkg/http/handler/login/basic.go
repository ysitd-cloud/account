package login

import (
	"database/sql"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/account/pkg/http/helper"
	"github.com/ysitd-cloud/account/pkg/http/middlewares"
	"github.com/ysitd-cloud/account/pkg/model/user"
	"github.com/ysitd-cloud/account/pkg/utils"
)

func basicForm(c *gin.Context) {
	session := middlewares.GetSession(c)
	nextUrl := c.DefaultQuery("next", "/")
	if !session.Exists("username") {
		helper.RenderAppView(c, http.StatusOK, "login", nil)
	} else {
		c.Redirect(http.StatusFound, nextUrl)
	}

}

func basicSubmit(c *gin.Context) {
	auth := false
	var reason string
	username := c.PostForm("username")
	password := c.PostForm("password")

	kernel := c.MustGet("kernel").(container.Kernel)
	pool := kernel.Make("db.pool").(utils.DatabasePool)

	instance, err := user.LoadFromDBWithUsername(pool, username)
	if instance == nil || err == sql.ErrNoRows {
		reason = "not_found"
	} else if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	} else if instance.ValidatePassword(password) {
		session := middlewares.GetSession(c)
		session.Set("username", instance.Username)
		session.Set("email", instance.Email)
		session.Set("avatar_url", instance.AvatarUrl)
		session.Set("display_name", instance.DisplayName)
		session.Save()
		auth = true
	} else {
		reason = "not_match"
	}

	next := c.DefaultPostForm("next", "/")

	if auth {
		c.Redirect(http.StatusFound, next)
	} else {
		redirect, _ := url.Parse("/login")
		query := redirect.Query()
		query.Set("next", next)
		query.Set("error", reason)
		redirect.RawQuery = query.Encode()
		c.Redirect(http.StatusFound, redirect.String())
	}

}
