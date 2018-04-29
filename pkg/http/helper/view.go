package helper

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"bytes"
	"code.ysitd.cloud/auth/account/pkg/http/middlewares"
	"code.ysitd.cloud/auth/account/pkg/model/user"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

var sidecarHost string

var client = &http.Client{
	Timeout: 10 * time.Second,
}

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

	sideCarURL.Path = fmt.Sprintf("/%s", view)
	sideCarURL.RawQuery = c.Request.URL.RawQuery

	body, err := json.Marshal(data)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	req, err := http.NewRequest("POST", sideCarURL.String(), bytes.NewReader(body))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		c.AbortWithError(http.StatusBadGateway, err)
	}

	defer res.Body.Close()

	c.Status(code)

	if _, err = io.Copy(c.Writer, res.Body); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Abort()
}
