package main

import (
	"os"
	"os/signal"

	"code.ysitd.cloud/component/account/pkg/bootstrap"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

func main() {
	if gin.IsDebugging() {
		logrus.SetLevel(logrus.DebugLevel)
	}
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go bootstrap.BootstrapPublicHTTPServer()
	go bootstrap.BootstrapGrpcProxy()
	go bootstrap.BootstrapGrpcServer()
	<-quit
}
