package storage

import (
	"fmt"

	"code.ysitd.cloud/component/account/pkg/utils"
	"github.com/RangelReale/osin"
	"github.com/garyburd/redigo/redis"
)

type Store struct {
	DB    utils.DatabasePool
	Redis *redis.Pool
}

func NewStore(dbPool utils.DatabasePool, redis *redis.Pool) osin.Storage {
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
