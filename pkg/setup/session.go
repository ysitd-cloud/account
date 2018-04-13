package setup

import (
	"code.ysitd.cloud/component/account/pkg/config"
	"code.ysitd.cloud/component/account/pkg/server/public/session"
	"github.com/facebookgo/inject"
)

func setupCoder(c *config.Config) *session.Coder {
	return session.NewTransCoder(c.Session.Key)
}

func injectCoder(c *config.Config, graph *inject.Graph) {
	coder := setupCoder(c)
	graph.Provide(&inject.Object{
		Value: coder,
	})
}
