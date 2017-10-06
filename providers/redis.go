package providers

import (
	"strconv"

	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/gin-utils/env"
)

type redisServiceProvider struct {
	*container.AbstractServiceProvider
}

func (*redisServiceProvider) Provides() []string {
	return []string{
		"redis.pool",
		"redis.address",
		"redis.db",
	}
}

func (*redisServiceProvider) Register(app container.Container) {
	app.Instance("redis.address", env.GetEnvWithDefault("REDIS_ADDRESS", "localhost:6379"))
	db, err := strconv.Atoi(env.GetEnvWithDefault("REDIS_DB", "0"))
	if err != nil {
		panic(err)
	}
	app.Instance("redis.db", db)
	app.Singleton("redis.pool", func(app container.Container) interface{} {
		address := app.Make("redis.address").(string)
		dbNo := app.Make("redis.db").(int)
		pool := redis.Pool{
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", address, redis.DialDatabase(dbNo))
			},
			TestOnBorrow: func(c redis.Conn, _ time.Time) (err error) {
				_, err = c.Do("PING")
				return
			},
		}

		return pool
	})
}
