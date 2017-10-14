package helper

import (
	"net/http"
	"net/url"
	"os"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang/groupcache/lru"
	"github.com/parnurzeal/gorequest"
)

var cache *lru.Cache = lru.New(8)
var sidecarHost string

type httpCache struct {
	etag    string
	content string
}

func init() {
	sidecarHost = os.Getenv("SIDECAR_URL")
}

func RenderAppView(c *gin.Context, code int, view string, data interface{}) {
	sideCarUrl, err := url.Parse(sidecarHost)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	req := gorequest.New()

	sideCarUrl.Path = fmt.Sprintf("/%s", view)
	sideCarUrl.RawQuery = c.Request.URL.RawQuery
	resultCaceh, exists := cache.Get(view)
	var pageCache httpCache
	if exists {
		pageCache = resultCaceh.(httpCache)
		req.Set("If-None-Match", pageCache.etag)
	}

	resp, body, errs := req.
		Post(sideCarUrl.String()).
		Send(data).
		End()
	if len(errs) != 0 {
		c.AbortWithError(http.StatusBadGateway, errs[0])
	}

	c.Status(code)
	if exists && resp.StatusCode == http.StatusNotModified {
		c.Writer.WriteString(pageCache.content)
	} else {
		c.Writer.WriteString(body)
		cache.Add(view, httpCache{
			etag:    resp.Header.Get("Etag"),
			content: body,
		})
	}
	c.Abort()
}
