package storage

import (
	"fmt"

	"github.com/RangelReale/osin"
	"github.com/gomodule/redigo/redis"
	"golang.ysitd.cloud/db"
)

type Store struct {
	DB    db.Opener
	Redis *redis.Pool
}

func NewStore(db db.Opener, redis *redis.Pool) osin.Storage {
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
