package main

import (
	"net"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/ysitd-cloud/account/pkg/http"
	"github.com/ysitd-cloud/account/pkg/providers"
	ginNet "github.com/ysitd-cloud/gin-utils/net"
	"google.golang.org/grpc"
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

	go app.Run(ginNet.GetAddress())

	server := providers.Kernel.Make("grpc.server").(*grpc.Server)
	listener := providers.Kernel.Make("grpc.listener").(net.Listener)
	server.Serve(listener)
}
