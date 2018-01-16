package metrics

import (
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ErrNotRegister     = errors.New("key not register")
	ErrNotRegisterRPC  = errors.Wrap(ErrNotRegister, "rpcEndpoints is not register")
	ErrNotRegisterHttp = errors.Wrap(ErrNotRegister, "httpEndpoints endpoint is not register")
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
	registry      registry
	rpc           *rpcCollector
	http          *rpcCollector
	rpcEndpoints  map[string]*rpcCollector
	httpEndpoints map[string]*rpcCollector
}

type rpcCollector struct {
	total *prometheus.CounterVec
	timer *prometheus.HistogramVec
}
