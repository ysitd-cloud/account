package providers

import (
	"database/sql"
	"os"

	"github.com/facebookgo/inject"
	_ "github.com/lib/pq"
	"github.com/tonyhhyip/go-di-container"
	"golang.ysitd.cloud/db"
)

func initDB() *db.GeneralOpener {
	return db.NewOpener("postgres", os.Getenv("DB_URL"))
}

func InjectDB(graph *inject.Graph) {
	graph.Provide(
		NewObject(initDB()),
	)
}

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
		return db.NewOpener(driver, url)
	})
}
