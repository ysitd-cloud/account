package storage

import (
	"database/sql"
	"fmt"

	"github.com/RangelReale/osin"
	"github.com/garyburd/redigo/redis"
)

type Store struct {
	DB    *sql.DB
	Redis *redis.Pool
}

func NewStore(db *sql.DB, pool *redis.Pool) osin.Storage {
	return &Store{
		db,
		pool,
	}
}

func (s *Store) Clone() osin.Storage {
	return s
}

func (s *Store) Close() {
	s.DB.Close()
	s.Redis.Close()
}

func makeKey(namespace, id string) string {
	return fmt.Sprintf("%s:%s", namespace, id)
}
