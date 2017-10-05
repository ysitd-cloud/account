package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/ysitd-cloud/account/http"
	"github.com/ysitd-cloud/gin-utils/net"
)

func main() {
	app := gin.Default()

	if gin.Mode() != gin.ReleaseMode {
		pprof.Register(app, nil)
	}

	http.Register(app)

	app.Run(net.GetAddress())
}
