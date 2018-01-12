package grpc

import (
	"code.ysitd.cloud/component/account/pkg/utils"
	"github.com/tonyhhyip/go-di-container"
)

type AccountService struct {
	Pool      utils.DatabasePool
	Container container.Container
}
