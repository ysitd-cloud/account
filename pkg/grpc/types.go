package grpc

import (
	"code.ysitd.cloud/common/go/db"
	"github.com/tonyhhyip/go-di-container"
)

type AccountService struct {
	Pool      db.Pool
	Container container.Container
}
