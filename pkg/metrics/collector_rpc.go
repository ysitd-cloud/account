package metrics

import (
	"time"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

func (c *collector) RegisterRPC(name string, labelsName []string) {
	labelsName = append(labelsName, "result")
	counter := newRPCCounter(name, labelsName)
	timer := newRPCTimer(name, labelsName)

	overall := newRPCCollector(counter, timer)
	logrus.WithFields(logrus.Fields{
		"target": "rpc",
		"name":   name,
	}).Debug("Register metrics collector")
	overall.register(c.registry)

	c.rpcEndpoints[name] = overall
}

func (c *collector) InvokeRPC(name string, labels prometheus.Labels) (finish chan<- bool, err error) {
	rpc, exists := c.rpcEndpoints[name]
	if !exists {
		return nil, errors.Wrapf(ErrNotRegisterRPC, "RPC %s is not registed", name)
	}

	channel := make(chan bool)
	go c.finishInvokeRPC(rpc, name, labels, channel)
	return channel, nil
}

func (c *collector) finishInvokeRPC(rpc *rpcCollector, endpoint string, labels prometheus.Labels, finish <-chan bool) {
	start := time.Now()
	overAllLabels := prometheus.Labels{
		"endpoint": endpoint,
	}
	if result := <-finish; result {
		labels["result"] = "success"
		overAllLabels["result"] = "success"
	} else {
		labels["result"] = "fail"
		overAllLabels["result"] = "fail"
	}
	duration := time.Now().Sub(start).Seconds()

	// Endpoint RPC Metrics
	rpc.total.With(labels).Inc()
	rpc.timer.With(labels).Observe(duration)

	// Overall RPC
	c.rpc.total.With(overAllLabels).Inc()
	c.rpc.timer.With(overAllLabels).Observe(duration)

	logger := logrus.WithField("target", "rpc")
	for k, v := range labels {
		logger = logger.WithField(k, v)
	}

	logger.Debug("Collect Metrics")
}
