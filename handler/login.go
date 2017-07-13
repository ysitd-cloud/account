package handler

import (
	"fmt"
	"net/http"
	"database/sql"
	"html/template"
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/ysitd-cloud/account/model"
	"github.com/ysitd-cloud/account/middlewares"
)

func LoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{
		"url": template.URL(fmt.Sprintf("/authorize?%s", c.Request.URL.RawQuery)),
	})
}

func LoginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	db := c.MustGet("db").(*sql.DB)

	user, err := model.LoadUserFromDBWithUsername(db, username)
	if err == sql.ErrNoRows {
		LoginForm(c)
		c.Abort()
		return
	}

	if user.ValidatePassword(password) {
		session := c.MustGet("sessions").(middlewares.Session)
		session.Set("username", user.Username)
		session.Set("email", user.Email)
		session.Set("avatar_url", user.AvatarUrl)
		session.Set("display_name", user.DisplayName)
		c.Next()
	} else {
		LoginForm(c)
		c.Abort()
	}
}
