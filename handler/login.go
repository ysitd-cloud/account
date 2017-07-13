package handler

import (
	"fmt"
	"net/http"
	"database/sql"
	"html/template"
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/ysitd-cloud/account/model"
	"github.com/ysitd-cloud/account/middlewares"
	"log"
)

func LoginForm(c *gin.Context) {
	session := middlewares.GetSession(c)
	if !session.Exists("username") {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{
			"url": template.URL(fmt.Sprintf("/authorize?%s", c.Request.URL.RawQuery)),
		})
		c.Abort()
	} else {
		c.Next()
	}

}

func LoginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	log.Printf("%s : %s", username, password)

	db := c.MustGet("db").(*sql.DB)

	user, err := model.LoadUserFromDBWithUsername(db, username)
	if err == sql.ErrNoRows {
		c.Next()
	} else if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	if user.ValidatePassword(password) {
		session := middlewares.GetSession(c)
		session.Set("username", user.Username)
		session.Set("email", user.Email)
		session.Set("avatar_url", user.AvatarUrl)
		session.Set("display_name", user.DisplayName)
		session.Save()
	}

	c.Next()
}
