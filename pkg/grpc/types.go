package grpc

import (
	"code.ysitd.cloud/auth/account/pkg/metrics"
	"code.ysitd.cloud/auth/account/third_party/forked/github.com/RangelReale/osin"
	"golang.ysitd.cloud/db"
)

type AccountService struct {
	Pool      *db.GeneralOpener  `inject:""`
	Server    *osin.Server       `inject:""`
	Collector *metrics.Collector `inject:""`
}
