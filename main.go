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
	bootstrap.BootstrapPublicHttpServer()
	bootstrap.BootstrapGrpcProxy()
	bootstrap.BootstrapGrpcServer()
	<-quit
}
