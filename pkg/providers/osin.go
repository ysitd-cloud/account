package providers

import (
	"code.ysitd.cloud/auth/account/pkg/storage"
	"code.ysitd.cloud/common/go/db"
	"github.com/RangelReale/osin"
	"github.com/gomodule/redigo/redis"
	"github.com/tonyhhyip/go-di-container"
)

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
		db := app.Make("db.pool").(db.Pool)
		redisDB := app.Make("redis.pool").(*redis.Pool)
		return storage.NewStore(db, redisDB)
	})

	app.Singleton("osin.config", func(app container.Container) interface{} {
		config := osin.NewServerConfig()
		config.AllowedAccessTypes = osin.AllowedAccessType{osin.AUTHORIZATION_CODE, osin.REFRESH_TOKEN}
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
