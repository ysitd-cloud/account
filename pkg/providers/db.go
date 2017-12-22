package providers

import (
	"os"

	"database/sql"

	_ "github.com/lib/pq"
	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/account/pkg/utils"
)

type databaseServiceProvider struct {
	*container.AbstractServiceProvider
}

func (*databaseServiceProvider) Provides() []string {
	return []string{
		"db",
		"db.pool",
		"db.postgres",
		"db.postgres.url",
	}
}

func (*databaseServiceProvider) Register(app container.Container) {
	app.Instance("db.postgres.url", os.Getenv("DB_URL"))
	app.Bind("db.postgres", func(app container.Container) interface{} {
		db, err := sql.Open("postgres", app.Make("db.postgres.url").(string))
		if err != nil {
			panic(err)
		}

		return db
	})
	app.Alias("db", "db.postgres")

	app.Singleton("db.pool", func(app container.Container) interface{} {
		driver := "postgres"
		url := app.Make("db.postgres.url").(string)
		return utils.NewDatabasePool(driver, url)
	})
}
