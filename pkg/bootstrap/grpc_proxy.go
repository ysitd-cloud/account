package bootstrap

import (
	"code.ysitd.cloud/component/account/pkg/metrics"
	"code.ysitd.cloud/gin/utils/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func BootstrapGrpcProxy() {
	proxy := Kernel.Make("grpc.proxy").(interfaces.Service)
	app := proxy.CreateService()
	app.GET("/metrics", bootstrapGrpcMetricsEndpoint())
	app.Run(":50050")
}

func bootstrapGrpcMetricsEndpoint() gin.HandlerFunc {
	collector := Kernel.Make("metrics").(metrics.Collector)
	handler := promhttp.HandlerFor(collector.GetGatherer(), promhttp.HandlerOpts{})
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
