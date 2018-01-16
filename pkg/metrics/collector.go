package metrics

import "github.com/prometheus/client_golang/prometheus"

func (c *collector) GetGatherer() prometheus.Gatherer {
	return c.registry
}

func (c *collector) init() {
	c.initHttpCollector()
	c.initRpcCollector()
	c.registry.MustRegister(c.rpc.timer, c.rpc.total, c.http.timer, c.http.total)
}

func (c *collector) initRpcCollector() {
	counter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Subsystem: "rpc",
		Name:      "count",
		Help:      "RPC total count",
		ConstLabels: prometheus.Labels{
			"component": "account",
		},
	}, []string{"result", "endpoint"})
	timer := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Subsystem: "rpc",
		Name:      "duration",
		Help:      "RPC call duration",
		ConstLabels: prometheus.Labels{
			"component": "account",
		},
	}, []string{"result", "endpoint"})
	c.rpc = newRPCCollector(counter, timer)
}

func (c *collector) initHttpCollector() {
	counter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Subsystem: "http",
		Name:      "count",
		Help:      "Http total count",
		ConstLabels: prometheus.Labels{
			"component": "account",
		},
	}, []string{"code", "endpoint"})
	timer := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Subsystem: "http",
		Name:      "duration",
		Help:      "Http call duration",
		ConstLabels: prometheus.Labels{
			"component": "account",
		},
	}, []string{"code", "endpoint"})
	c.http = newRPCCollector(counter, timer)
}
