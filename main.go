package main

import (
	"os"
	"os/signal"

	"code.ysitd.cloud/component/account/pkg/bootstrap"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go bootstrap.BootstrapPublicHttpServer()
	go bootstrap.BootstrapGrpcProxy()
	go bootstrap.BootstrapGrpcServer()
	<-quit
}
