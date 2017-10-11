package login

import (
	"database/sql"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/account/http/helper"
	"github.com/ysitd-cloud/account/http/middlewares"
	"github.com/ysitd-cloud/account/model"
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
	db := kernel.Make("db").(*sql.DB)
	defer db.Close()

	user, err := model.LoadUserFromDBWithUsername(db, username)
	if user == nil || err == sql.ErrNoRows {
		reason = "not_found"
	} else if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	} else if user.ValidatePassword(password) {
		session := middlewares.GetSession(c)
		session.Set("username", user.Username)
		session.Set("email", user.Email)
		session.Set("avatar_url", user.AvatarUrl)
		session.Set("display_name", user.DisplayName)
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
