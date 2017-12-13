package oauth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/account/pkg/http/middlewares"
)

func Callback(c *gin.Context) {
	session := middlewares.GetSession(c)
	usage, ok := session.Get("oauth:usage").(string)
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var prefix string = ""

	switch usage {
	case "connect":
		prefix = "/connect"
	case "login":
		prefix = "/login"
	default:
		// ignore
	}

	uri := fmt.Sprintf("%s/%s/callback?%s", prefix, c.Param("provider"), c.Request.URL.RawQuery)
	c.Redirect(http.StatusFound, uri)
	c.Abort()
}
