package utils

import "database/sql"

type DatabasePool interface {
	Acquire() (*sql.DB, error)
}

type databasePool struct {
	driver string
	url    string
}

func NewDatabasePool(driver, url string) DatabasePool {
	return &databasePool{
		driver: driver,
		url:    url,
	}
}

func (p *databasePool) Acquire() (*sql.DB, error) {
	return sql.Open(p.driver, p.url)
}
