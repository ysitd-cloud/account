package metrics

import (
	"time"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
)

func (c *collector) RegisterRPC(name string, labelsName []string) {
	counter := newRPCCounter(name, labelsName)
	timer := newRPCTimer(name, labelsName)

	overall := newRPCCollector(counter, timer)
	overall.register()

	c.rpc[name] = overall
}

func (c *collector) InvokeRPC(name string, labels prometheus.Labels) (finish chan<- bool, err error) {
	rpc, exists := c.rpc[name]
	if !exists {
		return nil, errors.Wrapf(ErrNotRegisterRPC, "RPC %s is not registed", name)
	}

	channel := make(chan bool)
	rpc.total.With(labels).Inc()
	go c.finishInvokeRPC(rpc.timer, labels, channel)
	return channel, nil
}

func (c *collector) finishInvokeRPC(counter *prometheus.HistogramVec, labels prometheus.Labels, finish <-chan bool) {
	start := time.Now()
	<-finish
	end := time.Now()
	duration := end.Sub(start)
	counter.With(labels).Observe(duration.Seconds())
}
