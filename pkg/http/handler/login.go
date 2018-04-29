package handler

import (
	"code.ysitd.cloud/auth/account/pkg/http/helper"
	"code.ysitd.cloud/auth/account/pkg/http/middlewares"
	"code.ysitd.cloud/auth/account/pkg/metrics"
	"code.ysitd.cloud/auth/account/pkg/model/user"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"golang.ysitd.cloud/db"
	"net/http"
	"net/url"
)

const (
	endpointLoginForm   = "login_form"
	endpointLoginSubmit = "login_submit"
)

type LoginHandler struct {
	Opener    *db.GeneralOpener  `inject:""`
	Collector *metrics.Collector `inject:""`
}

func (h *LoginHandler) RegisterRoutes(routes gin.IRoutes) {
	routes.GET("/login", h.loginForm)
	routes.POST("/login", h.loginPost)
}

func (h *LoginHandler) loginForm(c *gin.Context) {
	labels := prometheus.Labels{}
	finish, err := h.Collector.InvokeHTTP(endpointLoginForm, labels)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	defer func() {
		finish <- c.Writer.Status()
		close(finish)
	}()

	session := middlewares.GetSession(c)
	nextURL := c.DefaultQuery("next", "/")
	if !session.Exists("username") {
		helper.RenderAppView(c, http.StatusOK, "login", nil)
	} else {
		c.Redirect(http.StatusFound, nextURL)
	}
}

func (h *LoginHandler) loginPost(c *gin.Context) {
	auth := false
	var reason string
	username := c.PostForm("username")
	password := c.PostForm("password")

	labels := prometheus.Labels{
		"user": username,
	}

	finish, err := h.Collector.InvokeHTTP(endpointLoginSubmit, labels)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	defer func() {
		finish <- c.Writer.Status()
		close(finish)
	}()

	instance, err := user.LoadFromDBWithUsername(c.Request.Context(), h.Opener, username)
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
