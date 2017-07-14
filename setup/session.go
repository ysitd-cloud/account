package setup

import (
	sessions "github.com/ysitd-cloud/gin-sessions"
	"os"
)

func SetupSessionStore() (sessions.Store, error) {
	address := os.Getenv("REDIS_ADDRESS")
	if address == "" {
		address = "localhost:6379"
	}

	db := os.Getenv("REDIS_DB")
	if db == "" {
		db = "0"
	}

	secret := os.Getenv("SESSION_SECRET")
	return sessions.NewRedisStoreWithDB(1, "tcp", address, "", db, []byte(secret))
}
