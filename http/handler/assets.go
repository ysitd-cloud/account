package handler

import (
	"net/http"
	"net/http/httputil"

	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

var sidecarHost string
var sideCarUrl *url.URL

var assertProxy *httputil.ReverseProxy

func init() {
	sidecarHost = os.Getenv("SIDECAR_URL")
	sideCarUrl, _ = url.Parse(sidecarHost)
	assertProxy = &httputil.ReverseProxy{
		Director: assetsDirector,
	}
}

func assetsDirector(req *http.Request) {
	req.URL.Scheme = sideCarUrl.Scheme
	req.URL.Host = sideCarUrl.Host

	req.Header.Set("User-Agent", "Account Manager")
}

func AssetsProxy(c *gin.Context) {
	assertProxy.ServeHTTP(c.Writer, c.Request)
	c.Abort()
}
