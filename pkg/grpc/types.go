package grpc

import "database/sql"

type accountService struct {
	db *sql.DB
}
