package metrics

import "github.com/prometheus/client_golang/prometheus"

func (c *collector) GetGatherer() prometheus.Gatherer {
	return c.registry
}
