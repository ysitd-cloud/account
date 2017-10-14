package connect

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/account/http/helper"
	"github.com/ysitd-cloud/account/http/middlewares"
	"github.com/ysitd-cloud/account/model"
)

func listConnect(c *gin.Context) {
	session := middlewares.GetSession(c)
	user := model.User{
		Username:    session.Get("username").(string),
		DisplayName: session.Get("display_name").(string),
		Email:       session.Get("email").(string),
		AvatarUrl:   session.Get("avatar_url").(string),
	}

	vars := make(map[string]interface{})
	vars["user"] = user
	helper.RenderAppView(c, http.StatusOK, "connect", vars)
}
