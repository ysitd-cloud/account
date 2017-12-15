package grpc

import (
	"database/sql"

	"github.com/tonyhhyip/go-di-container"
)

type AccountService struct {
	DB        *sql.DB
	Container container.Container
}
