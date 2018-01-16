package bootstrap

import (
	"code.ysitd.cloud/gin/utils/interfaces"
	"code.ysitd.cloud/gin/utils/net"
)

func BootstrapPublicHttpServer() {
	service := Kernel.Make("http.service").(interfaces.Service)
	app := service.CreateService()
	go app.Run(net.GetAddress())
}
