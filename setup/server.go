package setup

import (
	"github.com/RangelReale/osin"
	"github.com/ysitd-cloud/account/storage"
)

func SetupOsinServer() *osin.Server {
	store := SetupOsinStore()

	config := SetupOsinServerConfig()

	return osin.NewServer(config, store)
}

func SetupOsinStore() osin.Storage {
	db, err := SetupDB()
	if err != nil {
		panic(err)
	}
	address := REDIS_ADDRESS
	if address == "" {
		address = "localhost:6379"
	}

	redisDB := REDIS_DB
	if redisDB == "" {
		redisDB = "0"
	}
	redis, err := SetupRedis(address, redisDB)
	if err != nil {
		panic(err)
	}

	return storage.NewStore(db, redis)
}

var config *osin.ServerConfig = nil

func SetupOsinServerConfig() *osin.ServerConfig {
	if config != nil {
		return config
	}
	config = osin.NewServerConfig()
	config.AllowedAccessTypes = osin.AllowedAccessType{osin.AUTHORIZATION_CODE, osin.REFRESH_TOKEN}
	config.AllowClientSecretInParams = true
	config.ErrorStatusCode = 400
}
