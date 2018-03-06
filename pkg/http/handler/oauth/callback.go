package oauth

import (
	"fmt"
	"net/http"

	"code.ysitd.cloud/component/account/pkg/http/middlewares"
	"github.com/gin-gonic/gin"
)

func Callback(c *gin.Context) {
	session := middlewares.GetSession(c)
	usage, ok := session.Get("oauth:usage").(string)
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	prefix := ""

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
