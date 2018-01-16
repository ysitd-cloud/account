package metrics

import (
	"time"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
)

func (c *collector) RegisterRPC(name string, labelsName []string) {
	labelsName = append(labelsName, "result")
	counter := newRPCCounter(name, labelsName)
	timer := newRPCTimer(name, labelsName)

	overall := newRPCCollector(counter, timer)
	overall.register(c.registry)

	c.rpc[name] = overall
}

func (c *collector) InvokeRPC(name string, labels prometheus.Labels) (finish chan<- bool, err error) {
	rpc, exists := c.rpc[name]
	if !exists {
		return nil, errors.Wrapf(ErrNotRegisterRPC, "RPC %s is not registed", name)
	}

	channel := make(chan bool)
	go c.finishInvokeRPC(rpc, labels, channel)
	return channel, nil
}

func (c *collector) finishInvokeRPC(rpc *rpcCollector, labels prometheus.Labels, finish <-chan bool) {
	start := time.Now()
	if result := <-finish; result {
		labels["result"] = "success"
	} else {
		labels["result"] = "failed"
	}
	duration := time.Now().Sub(start)
	rpc.total.With(labels).Inc()
	rpc.timer.With(labels).Observe(duration.Seconds())
}
