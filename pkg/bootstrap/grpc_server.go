package bootstrap

import (
	"net"

	"google.golang.org/grpc"
)

func BootstrapGrpcServer() {
	server := Kernel.Make("grpc.server").(*grpc.Server)
	listener := Kernel.Make("grpc.listener").(net.Listener)
	server.Serve(listener)
}
