package main

import (
	"net"

	proxy2 "code.ysitd.cloud/component/account/pkg/grpc/proxy"
	"code.ysitd.cloud/component/account/pkg/http"
	"code.ysitd.cloud/component/account/pkg/kernel"
	"code.ysitd.cloud/component/account/pkg/providers"
	ginNet "code.ysitd.cloud/gin/utils/net"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
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
