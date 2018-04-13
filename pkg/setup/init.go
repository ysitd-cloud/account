package setup

import (
	"os"

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

	if err := graph.Populate(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}
}
