package storage

import (
	"fmt"

	"code.ysitd.cloud/common/go/db"
	"github.com/RangelReale/osin"
	"github.com/garyburd/redigo/redis"
)

type Store struct {
	DB    db.Pool
	Redis *redis.Pool
}

func NewStore(dbPool db.Pool, redis *redis.Pool) osin.Storage {
	return &Store{
		dbPool,
		redis,
	}
}

func (s *Store) Clone() osin.Storage {
	return s
}

func (s *Store) Close() {}

func makeKey(namespace, id string) string {
	return fmt.Sprintf("%s:%s", namespace, id)
}
