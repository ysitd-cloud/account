package providers

import (
	"os"

	"github.com/facebookgo/inject"
	"github.com/gomodule/redigo/redis"
	"github.com/tonyhhyip/go-di-container"
	"golang.ysitd.cloud/gin/sessions"
	redisSession "golang.ysitd.cloud/gin/sessions/redis"
)

func initSession() *redisSession.RedisStore {
	secret := os.Getenv("SESSION_SECRET")
	store, err := redisSession.NewRedisStoreWithPool(nil, []byte(secret))
	if err != nil {
		panic(err)
	}
	return store
}

func InjectSession(graph *inject.Graph) {
	graph.Provide(
		NewObject(initSession()),
	)
}

type sessionServiceProvider struct {
	*container.AbstractServiceProvider
}

func (*sessionServiceProvider) Provides() []string {
	return []string{
		"session.secret",
		"session.store",
		"session.name",
		"session.middleware",
	}
}

func (*sessionServiceProvider) Register(app container.Container) {
	app.Instance("session.name", os.Getenv("SESSION_NAME"))
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
	app.Singleton("session.middleware", func(app container.Container) interface{} {
		name := app.Make("session.name").(string)
		store := app.Make("session.store").(sessions.Store)
		return sessions.Sessions(name, store, true)
	})
}
