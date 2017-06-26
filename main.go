package main

import (
	"os"
	"gopkg.in/oauth2.v3/manage"
	"github.com/go-oauth2/redis"
	"gopkg.in/oauth2.v3"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func createRedisTokenStore() (oauth2.TokenStore, error) {
	return redis.NewTokenStore(createRedisConfig())
}

func createRedisConfig() *redis.Config {
	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	return &redis.Config{
		Addr: host + ":" + port,
	}
}

func main() {
	db, err := gorm.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	manager := manage.NewDefaultManager()
	manager.MustTokenStorage(createRedisTokenStore())
}
