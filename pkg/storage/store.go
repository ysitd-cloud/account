package storage

import (
	"fmt"

	"code.ysitd.cloud/auth/account/third_party/forked/github.com/RangelReale/osin"
	"github.com/gomodule/redigo/redis"
	"golang.ysitd.cloud/db"
)

type Store struct {
	DB    *db.GeneralOpener
	Redis *redis.Pool
}

func NewStore(db *db.GeneralOpener, redis *redis.Pool) osin.Storage {
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
