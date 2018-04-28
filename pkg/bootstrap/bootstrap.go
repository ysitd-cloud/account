package bootstrap

import (
	"code.ysitd.cloud/auth/account/pkg/kernel"
	"code.ysitd.cloud/auth/account/pkg/providers"
)

// Kernel of everything
var Kernel = kernel.Kernel

func init() {
	providers.Register(Kernel)
}
