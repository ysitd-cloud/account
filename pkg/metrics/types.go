package metrics

import (
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ErrNotRegister     = errors.New("key not register")
	ErrNotRegisterRPC  = errors.Wrap(ErrNotRegister, "rpc is not register")
	ErrNotRegisterHttp = errors.Wrap(ErrNotRegister, "http endpoint is not register")
)

type registry interface {
	prometheus.Registerer
	prometheus.Gatherer
}

type Collector interface {
	GetGatherer() prometheus.Gatherer

	RegisterRPC(name string, labelsName []string)
	InvokeRPC(name string, labels prometheus.Labels) (finish chan<- bool, err error)

	RegisterHttp(endpoint string, labelsName []string)
	InvokeHttp(endpoint string, labels prometheus.Labels) (chan<- int, error)
}

type collector struct {
	registry registry
	rpc      map[string]*rpcCollector
	http     map[string]*rpcCollector
}

type rpcCollector struct {
	total *prometheus.CounterVec
	timer *prometheus.HistogramVec
}
