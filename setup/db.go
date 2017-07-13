package setup

import (
	_ "github.com/lib/pq"
	"database/sql"
	"os"
)

func SetupDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		return nil, err
	}

	return db, nil
}
