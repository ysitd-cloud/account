package providers

import (
	"database/sql"
	"net"
	"os"
	"path/filepath"

	"github.com/tonyhhyip/go-di-container"
	grpcService "github.com/ysitd-cloud/account/pkg/grpc"
	pb "github.com/ysitd-cloud/grpc-schema/account"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type grpcServiceProvder struct {
	*container.AbstractServiceProvider
}

func (*grpcServiceProvder) Provides() []string {
	return []string{
		"grpc.server",
		"grpc.listener",
	}
}

func (*grpcServiceProvder) Register(app container.Container) {
	app.Singleton("grpc.service", func(app container.Container) interface{} {
		db := app.Make("db").(*sql.DB)
		return &grpcService.AccountService{
			DB:        db,
			Container: app,
		}
	})

	app.Singleton("grpc.listener", func(app container.Container) interface{} {
		listener, err := net.Listen("tcp", "localhost:50051")
		if err != nil {
			panic(err)
		}

		return listener
	})

	app.Singleton("grpc.cert", func(app container.Container) interface{} {
		base := os.Getenv("CERT_PATH")
		certFile := filepath.Join(base, "tls.crt")
		keyFile := filepath.Join(base, "tls.key")
		cred, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			panic(err)
		}

		return cred
	})

	app.Singleton("grpc.server", func(app container.Container) interface{} {
		server := grpc.NewServer()

		service := app.Make("grpc.service").(pb.AccountServer)

		pb.RegisterAccountServer(server, service)

		return server
	})
}
