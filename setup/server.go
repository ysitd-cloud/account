package setup

import (
	"github.com/RangelReale/osin"
	"github.com/ysitd-cloud/account/storage"
)

func SetupOsinServer() (*osin.Server) {
	db, err := SetupDB()
	if err != nil {
		panic(err)
	}

	redis, err := SetupRedis()
	if err != nil {
		panic(err)
	}

	store := storage.NewStore(db, redis)

	config := osin.NewServerConfig()
	config.AllowedAccessTypes = osin.AllowedAccessType{osin.AUTHORIZATION_CODE, osin.REFRESH_TOKEN}
	config.AllowClientSecretInParams = true
	config.ErrorStatusCode = 400

	return osin.NewServer(config, store)
}
