package main

import (
	"code.ysitd.cloud/component/account/pkg/bootstrap"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	go bootstrap.BootstrapPublicHttpServer()
	go bootstrap.BootstrapGrpcProxy()
	bootstrap.BootstrapGrpcServer()
}
