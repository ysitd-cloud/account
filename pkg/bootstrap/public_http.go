package bootstrap

import (
	"code.ysitd.cloud/component/account/pkg/http"
	"code.ysitd.cloud/gin/utils/net"
	"github.com/gin-gonic/gin"
)

func BootstrapPublicHttpServer() {
	app := gin.Default()
	http.Register(app)
	app.Run(net.GetAddress())
}
