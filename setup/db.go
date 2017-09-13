package setup

import (
	"database/sql"
	"os"
	"strconv"

	"time"

	"github.com/garyburd/redigo/redis"
	_ "github.com/lib/pq"
)

func OpenDB() (*sql.DB, error) {
	return sql.Open("postgres", os.Getenv("DB_URL"))
}

func SetupDB() (*sql.DB, error) {
	db, err := OpenDB()
	if err != nil {
		return nil, err
	}

	for db.Ping() != nil {
		db2, err2 := OpenDB()
		db.Close()
		db = db2
		err = err2
	}

	return db, nil
}

var redisPool *redis.Pool = nil

func SetupRedis(address, db string) (*redis.Pool, error) {
	if redisPool != nil {
		return redisPool, nil
	}

	dbNo, err := strconv.Atoi(db)
	if err != nil {
		return nil, err
	}

	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address, redis.DialDatabase(dbNo))
		},
		TestOnBorrow: func(c redis.Conn, _ time.Time) (err error) {
			_, err = c.Do("PING")
			return
		},
	}

	redisPool = pool

	return pool, nil
}
