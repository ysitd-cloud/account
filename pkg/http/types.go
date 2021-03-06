package http

import (
	"code.ysitd.cloud/auth/account/pkg/metrics"
	"github.com/tonyhhyip/go-di-container"
)

type service struct {
	collector metrics.Collector
	app       container.Container
}
