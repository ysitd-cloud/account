package setup

import (
	sessions "github.com/ysitd-cloud/gin-sessions"
	"os"
)

var store sessions.Store = nil
var REDIS_ADDRESS string = os.Getenv("REDIS_ADDRESS")
var REDIS_DB string = os.Getenv("REDIS_DB")

func SetupSessionStore() (sessions.Store, error) {
	if store != nil {
		return store, nil
	}

	address := REDIS_ADDRESS
	if address == "" {
		address = "localhost:6379"
	}

	db := REDIS_DB
	if db == "" {
		db = "0"
	}
	pool, err := SetupRedis(address, db)
	if err != nil {
		return nil, err
	}
	secret := os.Getenv("SESSION_SECRET")
	store, err = sessions.NewRedisStoreWithPool(pool, []byte(secret))
	return store, err
}
