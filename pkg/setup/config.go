package setup

import (
	"github.com/facebookgo/inject"
	_ "github.com/joho/godotenv/autoload"

	"code.ysitd.cloud/component/account/pkg/config"
)

func setupConfig() *config.Config {
	c := config.NewConfigFromEnv()
	config.UpdateConfigWithFlag(c)
	return c
}

func injectConfig(c *config.Config, graph *inject.Graph) {
	graph.Provide(&inject.Object{
		Value: &c.Render.SideCarUrl,
		Name:  "sidecar",
	})
}
