package setup

import (
	"code.ysitd.cloud/component/account/pkg/config"
	"github.com/facebookgo/inject"
)

type configFun func(c *config.Config, graph *inject.Graph)

func init() {
	c := setupConfig()
	logger := setupLogger(c)

	graph := &inject.Graph{
		Logger: logger,
	}

	graph.Provide(
		&inject.Object{Value: &publicService},
	)

	funs := []configFun{
		injectConfig,
		injectDB,
		injectCoder,
		injectHttpClient,
	}

	for _, fun := range funs {
		fun(c, graph)
	}
}
