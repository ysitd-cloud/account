package setup

import (
	"code.ysitd.cloud/component/account/pkg/config"
	"github.com/sirupsen/logrus"
)

func setupLogger(c *config.Config) (l *logrus.Logger) {
	l = logrus.New()
	if c.Verbose {
		l.SetLevel(logrus.DebugLevel)
	}
	return
}
