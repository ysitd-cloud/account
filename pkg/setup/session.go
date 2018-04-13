package setup

import (
	"code.ysitd.cloud/component/account/pkg/config"
	"code.ysitd.cloud/component/account/pkg/server/public/session"
)

func setupCoder(c *config.Config) *session.Coder {
	return session.NewTransCoder(c.Session.Key)
}
