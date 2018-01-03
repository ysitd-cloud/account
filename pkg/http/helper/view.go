package helper

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	"github.com/ysitd-cloud/account/pkg/http/middlewares"
	"github.com/ysitd-cloud/account/pkg/model/user"
)

var sidecarHost string

func init() {
	sidecarHost = os.Getenv("SIDECAR_URL")
}

func RenderAppView(c *gin.Context, code int, view string, data map[string]interface{}) {

	if session := middlewares.GetSession(c); session.Exists("username") {
		if data == nil {
			data = make(map[string]interface{})
		}
		instance := user.User{
			Username:    session.Get("username").(string),
			DisplayName: session.Get("display_name").(string),
			Email:       session.Get("email").(string),
			AvatarUrl:   session.Get("avatar_url").(string),
		}
		data["user"] = instance
	}

	sideCarUrl, err := url.Parse(sidecarHost)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	req := gorequest.New()

	sideCarUrl.Path = fmt.Sprintf("/%s", view)
	sideCarUrl.RawQuery = c.Request.URL.RawQuery

	_, body, errs := req.
		Post(sideCarUrl.String()).
		Send(data).
		End()
	if len(errs) != 0 {
		c.AbortWithError(http.StatusBadGateway, errs[0])
	}

	c.Status(code)
	c.Writer.WriteString(body)
	c.Abort()
}
