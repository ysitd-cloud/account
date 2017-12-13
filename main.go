package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/ysitd-cloud/account/pkg/http"
	"github.com/ysitd-cloud/account/pkg/providers"
	"github.com/ysitd-cloud/gin-utils/net"
)

func init() {
	app := gin.Default()
	providers.Kernel.Instance("http.server", app)
}

func main() {
	app := providers.Kernel.Make("http.server").(*gin.Engine)

	if gin.Mode() != gin.ReleaseMode {
		pprof.Register(app, nil)
	}

	http.Register(app)

	app.Run(net.GetAddress())
}
