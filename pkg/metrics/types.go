package metrics

import (
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ErrNotRegister     = errors.New("key not register")
	ErrNotRegisterRPC  = errors.Wrap(ErrNotRegister, "rpcEndpoints is not register")
	ErrNotRegisterHTTP = errors.Wrap(ErrNotRegister, "httpEndpoints endpoint is not register")
)

type DoneFunc func(result bool)

type registry interface {
	prometheus.Registerer
	prometheus.Gatherer
}

type Collector struct {
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
