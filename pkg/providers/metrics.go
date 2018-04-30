package providers

import (
	"code.ysitd.cloud/auth/account/pkg/metrics"
	"github.com/facebookgo/inject"
	"github.com/prometheus/client_golang/prometheus"
)

func initMetrics() *metrics.Collector {
	return metrics.NewCollector(prometheus.DefaultRegisterer.(*prometheus.Registry))
}

func InjectMetrics(graph *inject.Graph) {
	graph.Provide(
		NewObject(initMetrics()),
	)
}
