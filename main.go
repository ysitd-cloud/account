package main

import (
	"net"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	proxy2 "github.com/ysitd-cloud/account/pkg/grpc/proxy"
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

	proxy := kernel.Kernel.Make("grpc.proxy").(proxy2.GrpcProxy)
	engine := proxy.CreateApp()
	go engine.Run(":50050")

	server := kernel.Kernel.Make("grpc.server").(*grpc.Server)
	listener := kernel.Kernel.Make("grpc.listener").(net.Listener)
	server.Serve(listener)
}
