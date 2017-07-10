package setup

import (
	_ "github.com/lib/pq"
	"github.com/RangelReale/osin"
	"database/sql"
	"os"
	"github.com/ory/osin-storage/storage/postgres"
)

func SetupOsinServer() (*osin.Server) {
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}

	store := postgres.New(db)
	config := osin.NewServerConfig()
	config.AllowedAccessTypes = osin.AllowedAccessType{osin.AUTHORIZATION_CODE, osin.REFRESH_TOKEN}
	config.ErrorStatusCode = 400
	return osin.NewServer(config, store)
}
