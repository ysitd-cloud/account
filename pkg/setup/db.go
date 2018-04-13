package setup

import (
	_ "github.com/lib/pq"

	"code.ysitd.cloud/common/go/db"
	"code.ysitd.cloud/component/account/pkg/config"
)

func setupDatabase(c *config.Config) db.Pool {
	return db.NewPool(c.Database.Driver, c.Database.Driver)
}
