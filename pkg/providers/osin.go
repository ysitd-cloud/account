package providers

import (
	"code.ysitd.cloud/auth/account/pkg/storage"
	"code.ysitd.cloud/auth/account/third_party/forked/github.com/RangelReale/osin"
	"github.com/facebookgo/inject"
	"github.com/gomodule/redigo/redis"
	"github.com/tonyhhyip/go-di-container"
	"golang.ysitd.cloud/db"
	"time"
)

func initOsinConfig() (config *osin.ServerConfig) {
	config = osin.NewServerConfig()
	config.AllowedAccessTypes = osin.AllowedAccessType{osin.AuthorizationCode, osin.RefreshToken}
	config.AllowClientSecretInParams = true
	config.ErrorStatusCode = 400
	return
}

func InjectOsin(graph *inject.Graph) {
	graph.Provide(
		NewObject(initOsinConfig()),
		NewNamedObject("osin storage", new(storage.Store)),
		NewNamedObject("osin authorize token gen", new(osin.AuthorizeTokenGenDefault)),
		NewNamedObject("osin access token gen", new(osin.AccessTokenGenDefault)),
		NewNamedObject("osin now func", func() time.Time { return time.Now().UTC() }),
	)
}

type osinServiceProvider struct {
	*container.AbstractServiceProvider
}

func (*osinServiceProvider) Provides() []string {
	return []string{
		"osin.storage",
		"osin.config",
		"osin.server",
	}
}

func (*osinServiceProvider) Register(app container.Container) {
	app.Bind("osin.storage", func(app container.Container) interface{} {
		db := app.Make("db.pool").(*db.GeneralOpener)
		redisDB := app.Make("redis.pool").(*redis.Pool)
		return storage.NewStore(db, redisDB)
	})

	app.Singleton("osin.config", func(app container.Container) interface{} {
		config := osin.NewServerConfig()
		config.AllowedAccessTypes = osin.AllowedAccessType{osin.AuthorizationCode, osin.RefreshToken}
		config.AllowClientSecretInParams = true
		config.ErrorStatusCode = 400
		return config
	})

	app.Bind("osin.server", func(app container.Container) interface{} {
		store := app.Make("osin.storage").(osin.Storage)
		config := app.Make("osin.config").(*osin.ServerConfig)

		return osin.NewServer(config, store)
	})
}
