package metrics

import "github.com/prometheus/client_golang/prometheus"

func (c *rpcCollector) register(registerer prometheus.Registerer) {
	registerer.MustRegister(c.total, c.timer)
}
