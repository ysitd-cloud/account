package bootstrap

import (
	"code.ysitd.cloud/gin/utils/interfaces"
	"code.ysitd.cloud/gin/utils/net"
)

func BootstrapPublicHTTPServer() {
	service := Kernel.Make("http.service").(interfaces.Service)
	app := service.CreateService()
	app.Run(net.GetAddress())
}
