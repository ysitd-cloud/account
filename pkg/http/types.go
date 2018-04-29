package http

import (
	"code.ysitd.cloud/auth/account/pkg/http/handler"
	"code.ysitd.cloud/auth/account/pkg/metrics"
	"github.com/tonyhhyip/go-di-container"
)

type service struct {
	Collector    *metrics.Collector `inject:""`
	app          container.Container
	LoginHandler *handler.LoginHandler `inject:""`
}
