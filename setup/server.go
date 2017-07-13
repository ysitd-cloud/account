package setup

import (
	"github.com/RangelReale/osin"
	"github.com/ory/osin-storage/storage/postgres"
)

func SetupOsinServer() (*osin.Server) {
	db, err := SetupDB()
	if err != nil {
		panic(err)
	}

	store := postgres.New(db)
	config := osin.NewServerConfig()
	config.AllowedAccessTypes = osin.AllowedAccessType{osin.AUTHORIZATION_CODE, osin.REFRESH_TOKEN}
	config.ErrorStatusCode = 400
	return osin.NewServer(config, store)
}
