package bootstrap

import (
	"code.ysitd.cloud/component/account/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/tonyhhyip/vodka"
)

// BootstrapGrpcProxy start running grpc proxy service
func BootstrapGrpcProxy() {
	router := Kernel.Make("grpc.proxy").(*vodka.Router)
	logger := Kernel.Make("logger").(*logrus.Logger)
	server := vodka.New(":50050")
	server.SetLogger(logger.WithField("source", "grpc.proxy"))
	router.GET("/metrics", bootstrapMetricsEndpoint())

	server.ListenAndServe(router.Handler())
}

func bootstrapMetricsEndpoint() vodka.HandlerFunc {
	collector := Kernel.Make("metrics").(metrics.Collector)
	handler := promhttp.HandlerFor(collector.GetGatherer(), promhttp.HandlerOpts{})
	return vodka.HandlerFunc(func(c *vodka.Context) {
		handler.ServeHTTP(c.Response, c.Request)
	})
}
