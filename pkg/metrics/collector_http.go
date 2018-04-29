package metrics

import (
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

func (c *Collector) RegisterHTTP(endpoint string, labelsName []string) {
	labelsName = append(labelsName, "code")
	counter := newHTTPCounter(endpoint, labelsName)
	timer := newHTTPTimer(endpoint, labelsName)

	rpc := newRPCCollector(counter, timer)
	logrus.WithFields(logrus.Fields{
		"target":   "http",
		"endpoint": endpoint,
		"labels":   labelsName,
	}).Debug("Register metrics Collector")
	rpc.register(c.registry)

	c.httpEndpoints[endpoint] = rpc
}

func (c *Collector) InvokeHTTP(endpoint string, labels prometheus.Labels) (done HTTPDoneFunc, err error) {
	rpc, exists := c.httpEndpoints[endpoint]
	if !exists {
		return nil, errors.Wrapf(ErrNotRegisterHTTP, "Endpoint %s is not register", endpoint)
	}
	done = c.wrapFinishInvokeHTTP(rpc, endpoint, labels)
	return
}

func (c *Collector) wrapFinishInvokeHTTP(rpc *rpcCollector, endpoint string, labels prometheus.Labels) HTTPDoneFunc {
	start := time.Now()

	return func(code int) {
		duration := time.Now().Sub(start).Seconds()

		statusCode := strconv.FormatInt(int64(code), 10)

		labels["code"] = statusCode
		rpc.timer.With(labels).Observe(duration)
		rpc.total.With(labels).Inc()

		overAllLabels := prometheus.Labels{
			"code":     statusCode,
			"endpoint": endpoint,
		}
		c.http.total.With(overAllLabels).Inc()
		c.http.timer.With(overAllLabels).Observe(duration)

		loggerFields := make(map[string]interface{})
		loggerFields["target"] = "http"
		for k, v := range labels {
			loggerFields[k] = v
		}
		logger := logrus.WithField("target", "http")

		// Endpoint RPC Metrics
		rpc.total.With(labels).Inc()
		rpc.timer.With(labels).Observe(duration)

		logger.Debug("Collect Metrics")
	}
}
