package grpc

import (
	"code.ysitd.cloud/common/go/db"
	"code.ysitd.cloud/component/account/pkg/metrics"
	"github.com/tonyhhyip/go-di-container"
)

type AccountService struct {
	Pool      db.Pool
	Container container.Container
	Collector metrics.Collector
}
