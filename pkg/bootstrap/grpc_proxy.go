package bootstrap

import (
	"code.ysitd.cloud/auth/account/pkg/grpc/proxy"
	"code.ysitd.cloud/auth/account/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

// RunGrpcProxy start running grpc proxy service
func RunGrpcProxy() {
	handler := Kernel.Make("grpc.proxy").(*proxy.Handler)
	handler.ConfigMetrics(bootstrapMetricsEndpoint())
	http.ListenAndServe(":50050", handler)
}

func bootstrapMetricsEndpoint() http.Handler {
	collector := Kernel.Make("metrics").(metrics.Collector)
	handler := promhttp.HandlerFor(collector.GetGatherer(), promhttp.HandlerOpts{})
	return handler
}
