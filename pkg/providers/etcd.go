package providers

import (
	"os"

	"github.com/tonyhhyip/go-di-container"
)

type etcdServiceProvider struct {
	*container.AbstractServiceProvider
}

func (*etcdServiceProvider) Provides() []string {
	return []string{
		"etcd.host",
	}
}

func (*etcdServiceProvider) Register(app container.Container) {
	app.Instance("etcd.host", os.Getenv("ETCD_HOST"))
}
