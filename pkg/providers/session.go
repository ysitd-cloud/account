package providers

import (
	"os"

	"code.ysitd.cloud/gin/sessions"
	"code.ysitd.cloud/gin/sessions/etcd"
	"github.com/tonyhhyip/go-di-container"
)

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
	app.Instance("session.path", os.Getenv("SESSION_PATH"))
	app.Singleton("session.store", func(app container.Container) interface{} {
		secret := app.Make("session.secret").(string)
		host := app.Make("etcd.host").(string)
		path := app.Make("session.path").(string)
		store, err := etcd.NewEtcdStoreFromAddr([]string{host}, path, []byte(secret))
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
