package providers

import (
	"os"

	"github.com/garyburd/redigo/redis"
	"github.com/tonyhhyip/go-di-container"
	redisSession "github.com/ysitd-cloud/gin-sessions/redis"
)

type sessionServiceProvider struct {
	*container.AbstractServiceProvider
}

func (*sessionServiceProvider) Provides() []string {
	return []string{
		"session.secret",
		"session.store",
	}
}

func (*sessionServiceProvider) Register(app container.Container) {
	app.Instance("session.secret", os.Getenv("SESSION_SECRET"))
	app.Singleton("session.store", func(app container.Container) interface{} {
		secret := app.Make("session.secret").(string)
		pool := app.Make("redis.pool").(*redis.Pool)
		store, err := redisSession.NewRedisStoreWithPool(pool, []byte(secret))
		if err != nil {
			panic(err)
		}
		return store
	})
}
