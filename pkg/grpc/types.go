package grpc

import (
	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/account/pkg/utils"
)

type AccountService struct {
	Pool      utils.DatabasePool
	Container container.Container
}
