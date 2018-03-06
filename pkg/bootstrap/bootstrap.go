package bootstrap

import (
	"code.ysitd.cloud/component/account/pkg/kernel"
	"code.ysitd.cloud/component/account/pkg/providers"
)

// Kernel of everything
var Kernel = kernel.Kernel

func init() {
	providers.Register(Kernel)
}
