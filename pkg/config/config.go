package config

import "code.ysitd.cloud/component/account/pkg/config/env"

type Config struct {
	Verbose  bool
	Database *database
}

type database struct {
	Driver     string
	DataSource string `env:"DATABASE_URL"`
}

func newDatabaseFromEnv() *database {
	db := &database{
		Driver: "postgres",
	}
	env.InjectWithEnv(db)
	return db
}
