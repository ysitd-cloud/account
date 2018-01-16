package main

import (
	"net"

	proxy2 "code.ysitd.cloud/component/account/pkg/grpc/proxy"
	"code.ysitd.cloud/component/account/pkg/http"
	"code.ysitd.cloud/component/account/pkg/kernel"
	"code.ysitd.cloud/component/account/pkg/providers"
	ginNet "code.ysitd.cloud/gin/utils/net"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

func init() {
	app := gin.Default()
	kernel.Kernel.Instance("http.server", app)
	providers.Register(kernel.Kernel)
}

func main() {
	{
		app := kernel.Kernel.Make("http.server").(*gin.Engine)
		http.Register(app)
		go app.Run(ginNet.GetAddress())
	}

	{
		proxy := kernel.Kernel.Make("grpc.proxy").(proxy2.GrpcProxy)
		app := proxy.CreateApp()
		handler := promhttp.Handler()
		app.GET("/metrics", func(c *gin.Context) {
			handler.ServeHTTP(c.Writer, c.Request)
		})
		go app.Run(":50050")
	}

	{
		server := kernel.Kernel.Make("grpc.server").(*grpc.Server)
		listener := kernel.Kernel.Make("grpc.listener").(net.Listener)
		server.Serve(listener)
	}
}
