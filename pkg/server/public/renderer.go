package public

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/tonyhhyip/vodka"
)

type Renderer struct {
	Client      *http.Client `inject:""`
	SideCarHost string       `inject:"sidecar"`
}

func (r *Renderer) Render(c *vodka.Context, code int, view string, data map[string]interface{}) {
	sideCarUrl := &url.URL{
		Scheme: "http",
		Host:   r.SideCarHost,
	}

	sideCarUrl.Path = "/" + view
	sideCarUrl.RawQuery = c.Request.URL.RawQuery

	var body bytes.Buffer

	b, err := json.Marshal(data)
	if err != nil {
		c.Logger().Error(err)
		c.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	body.Write(b)

	req, err := http.NewRequest("POST", sideCarUrl.String(), &body)

	if err != nil {
		c.Logger().Error(err)
		c.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := r.Client.Do(req)

	if err != nil {
		c.Logger().Error(err)
		c.Error(err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	c.Status(code)

	io.Copy(c.Response, resp.Body)
}
