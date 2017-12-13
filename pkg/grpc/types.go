package grpc

import "database/sql"

type AccountService struct {
	DB *sql.DB
}
