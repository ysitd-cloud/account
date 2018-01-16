package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tonyhhyip/go-di-container"
)

func NewServiceProvider(app container.Container) container.ServiceProvider {
	sp := &serviceProvider{
		AbstractServiceProvider: container.NewAbstractServiceProvider(true),
	}

	sp.SetContainer(app)

	return sp
}

func NewCollector(registry registry) Collector {
	c := &collector{
		rpcEndpoints:  make(map[string]*rpcCollector),
		httpEndpoints: make(map[string]*rpcCollector),
		registry:      registry,
	}

	c.init()

	return c
}

func newRPCCollector(total *prometheus.CounterVec, timer *prometheus.HistogramVec) *rpcCollector {
	return &rpcCollector{
		total: total,
		timer: timer,
	}
}

func newRPCCounter(name string, labelsName []string) *prometheus.CounterVec {
	return prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "rpc",
		Subsystem: name,
		Name:      "requests_total",
		Help:      "RPC call count for " + name,
	}, labelsName)
}

func newRPCTimer(name string, labelsName []string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "rpc",
		Subsystem: name,
		Name:      "duration",
		Help:      "RPC call duration for " + name,
	}, labelsName)
}

func newHttpCounter(endpoint string, labelsName []string) *prometheus.CounterVec {
	return prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "http",
		Subsystem: endpoint,
		Name:      "requests_total",
		Help:      "Http call count for " + endpoint,
	}, labelsName)
}

func newHttpTimer(endpoint string, labelsName []string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "http",
		Subsystem: endpoint,
		Name:      "duration",
		Help:      "Http call duration for " + endpoint,
	}, labelsName)
}
