package metrics

import (
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ErrNotRegister    = errors.New("key not register")
	ErrNotRegisterRPC = errors.Wrap(ErrNotRegister, "rpc is not register")
)

type Collector interface {
	RegisterRPC(name string, labelsName []string)
	InvokeRPC(name string, labels prometheus.Labels) (finish chan<- bool, err error)
}

type collector struct {
	rpc map[string]*rpcCollector
}

type rpcCollector struct {
	total *prometheus.CounterVec
	timer *prometheus.HistogramVec
}
