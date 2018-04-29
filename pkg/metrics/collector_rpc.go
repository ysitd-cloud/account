package metrics

import (
	"time"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

func (c *Collector) RegisterRPC(name string, labelsName []string) {
	labelsName = append(labelsName, "result")
	counter := newRPCCounter(name, labelsName)
	timer := newRPCTimer(name, labelsName)

	overall := newRPCCollector(counter, timer)
	logrus.WithFields(logrus.Fields{
		"target": "rpc",
		"name":   name,
	}).Debug("Register metrics Collector")
	overall.register(c.registry)

	c.rpcEndpoints[name] = overall
}

func (c *Collector) InvokeRPC(name string, labels prometheus.Labels) (done RPCDoneFunc, err error) {
	rpc, exists := c.rpcEndpoints[name]
	if !exists {
		return nil, errors.Wrapf(ErrNotRegisterRPC, "RPC %s is not registed", name)
	}

	done = c.wrapFinishInvokeRPC(rpc, name, labels)
	return
}

func (c *Collector) wrapFinishInvokeRPC(rpc *rpcCollector, endpoint string, labels prometheus.Labels) RPCDoneFunc {
	start := time.Now()
	overAllLabels := prometheus.Labels{
		"endpoint": endpoint,
	}

	return func(result bool) {
		if result {
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

		loggerFields := make(map[string]interface{})
		loggerFields["target"] = "rpc"
		for k, v := range labels {
			loggerFields[k] = v
		}

		logger := logrus.WithFields(loggerFields)

		logger.Debug("Collect Metrics")
	}
}
