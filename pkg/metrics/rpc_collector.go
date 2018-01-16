package metrics

import "github.com/prometheus/client_golang/prometheus"

func (c *rpcCollector) register() {
	prometheus.MustRegister(c.total, c.timer)
}
