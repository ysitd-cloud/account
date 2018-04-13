package config

import "code.ysitd.cloud/component/account/pkg/config/env"

type Config struct {
	Verbose  bool
	Database *database
	Render   *render
}

type render struct {
	SideCarUrl string `env:"SIDECAR_URL"`
}

func newRenderFromEnv() (r *render) {
	r = new(render)
	env.InjectWithEnv(r)
	return
}

type database struct {
	Driver     string
	DataSource string `env:"DATABASE_URL"`
}

func newDatabaseFromEnv() (db *database) {
	db = &database{
		Driver: "postgres",
	}
	env.InjectWithEnv(db)
	return
}
