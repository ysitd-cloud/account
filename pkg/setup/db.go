package setup

import (
	"github.com/facebookgo/inject"
	_ "github.com/lib/pq"

	"code.ysitd.cloud/common/go/db"
	"code.ysitd.cloud/component/account/pkg/config"
)

func setupDatabase(c *config.Config) db.Pool {
	return db.NewPool(c.Database.Driver, c.Database.Driver)
}

func injectDB(c *config.Config, graph *inject.Graph) {
	pool := setupDatabase(c)
	graph.Provide(&inject.Object{
		Value: pool,
		Name:  "db",
	})
}
