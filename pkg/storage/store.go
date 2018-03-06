package storage

import (
	"fmt"

	"code.ysitd.cloud/common/go/db"
	"github.com/RangelReale/osin"
	"github.com/garyburd/redigo/redis"
)

type Store struct {
	DB    db.DBOpener
	Redis *redis.Pool
}

func NewStore(db db.DBOpener, redis *redis.Pool) osin.Storage {
	return &Store{
		db,
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
