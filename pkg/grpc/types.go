package grpc

import (
	"code.ysitd.cloud/auth/account/pkg/metrics"
	"github.com/tonyhhyip/go-di-container"
	"golang.ysitd.cloud/db"
)

type AccountService struct {
	Pool      *db.GeneralOpener `inject:""`
	Container container.Container
	Collector *metrics.Collector `inject:""`
}
