package setup

import (
	_ "github.com/joho/godotenv/autoload"

	"code.ysitd.cloud/component/account/pkg/config"
)

func setupConfig() *config.Config {
	c := config.NewConfigFromEnv()
	config.UpdateConfigWithFlag(c)
	return c
}
