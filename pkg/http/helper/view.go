package helper

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"code.ysitd.cloud/component/account/pkg/http/middlewares"
	"code.ysitd.cloud/component/account/pkg/model/user"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
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

	sideCarURL, err := url.Parse(sidecarHost)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	req := gorequest.New()

	sideCarURL.Path = fmt.Sprintf("/%s", view)
	sideCarURL.RawQuery = c.Request.URL.RawQuery

	_, body, errs := req.
		Post(sideCarURL.String()).
		Send(data).
		End()
	if len(errs) != 0 {
		c.AbortWithError(http.StatusBadGateway, errs[0])
	}

	c.Status(code)
	c.Writer.WriteString(body)
	c.Abort()
}
