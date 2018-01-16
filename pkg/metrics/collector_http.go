package metrics

import (
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
)

func (c *collector) RegisterHttp(endpoint string, labelsName []string) {
	counter := newHttpCounter(endpoint, labelsName)
	timer := newHttpTimer(endpoint, labelsName)

	rpc := newRPCCollector(counter, timer)
	rpc.register(c.registry)

	c.http[endpoint] = rpc
}

func (c *collector) InvokeHttp(endpoint string, labels prometheus.Labels) (chan<- int, error) {
	rpc, exists := c.http[endpoint]
	if !exists {
		return nil, errors.Wrapf(ErrNotRegisterHttp, "Endpoint %s is not register", endpoint)
	}
	channel := make(chan int)
	go c.finishHttp(endpoint, labels, rpc, channel)
	return channel, nil
}

func (c *collector) finishHttp(
	endpoint string,
	labels prometheus.Labels,
	rpc *rpcCollector,
	channel <-chan int,
) {
	start := time.Now()
	code := <-channel
	duration := time.Now().Sub(start).Seconds()
	labels["code"] = strconv.FormatInt(int64(code), 10)
	rpc.timer.With(labels).Observe(duration)
	rpc.total.With(labels).Inc()
}
