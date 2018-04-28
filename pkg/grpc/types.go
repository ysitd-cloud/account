package grpc

import (
	"code.ysitd.cloud/auth/account/pkg/metrics"
	"github.com/tonyhhyip/go-di-container"
	"golang.ysitd.cloud/db"
)

type AccountService struct {
	Pool      db.Opener
	Container container.Container
	Collector metrics.Collector
}
