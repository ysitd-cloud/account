package setup

import (
	"database/sql"
	"os"
	"strconv"

	"github.com/garyburd/redigo/redis"
	_ "github.com/lib/pq"
)

func SetupDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func SetupRedis() (*redis.Pool, error) {
	address := os.Getenv("REDIS_ADDRESS")
	if address == "" {
		address = "localhost:6379"
	}

	db := os.Getenv("REDIS_DB")
	if db == "" {
		db = "0"
	}

	dbNo, err := strconv.Atoi(db)
	if err != nil {
		return nil, err
	}

	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address, redis.DialDatabase(dbNo))
		},
	}

	return pool, nil
}
