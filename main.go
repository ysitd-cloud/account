package main

import (
	"net"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/ysitd-cloud/account/pkg/http"
	"github.com/ysitd-cloud/account/pkg/kernel"
	"github.com/ysitd-cloud/account/pkg/providers"
	ginNet "github.com/ysitd-cloud/gin-utils/net"
	"google.golang.org/grpc"
)

func init() {
	app := gin.Default()
	kernel.Kernel.Instance("http.server", app)
	providers.Register(kernel.Kernel)
}

func main() {
	app := kernel.Kernel.Make("http.server").(*gin.Engine)

	if gin.Mode() != gin.ReleaseMode {
		pprof.Register(app, nil)
	}

	http.Register(app)

	go app.Run(ginNet.GetAddress())

	server := kernel.Kernel.Make("grpc.server").(*grpc.Server)
	listener := kernel.Kernel.Make("grpc.listener").(net.Listener)
	server.Serve(listener)
}
