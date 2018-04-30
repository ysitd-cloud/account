package storage

import (
	"code.ysitd.cloud/auth/account/third_party/forked/github.com/RangelReale/osin"
	"github.com/gomodule/redigo/redis"
)

func (s *Store) SaveAuthorize(data *osin.AuthorizeData) (err error) {
	conn := s.Redis.Get()
	if err := conn.Err(); err != nil {
		return err
	}

	defer conn.Close()

	payload, err := encode(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SETEX", makeKey("auth", data.Code), data.ExpiresIn, string(payload))
	return err
}

func (s *Store) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	conn := s.Redis.Get()
	if err := conn.Err(); err != nil {
		return nil, err
	}

	defer conn.Close()

	var (
		rawAuthGob interface{}
		err        error
	)

	if rawAuthGob, err = conn.Do("GET", makeKey("auth", code)); err != nil {
		return nil, err
	}
	if rawAuthGob == nil {
		return nil, nil
	}

	authGob, _ := redis.Bytes(rawAuthGob, err)

	var auth osin.AuthorizeData
	err = decode(authGob, &auth)
	return &auth, err
}

func (s *Store) RemoveAuthorize(code string) (err error) {
	conn := s.Redis.Get()
	if err := conn.Err(); err != nil {
		return err
	}

	defer conn.Close()

	_, err = conn.Do("DEL", makeKey("auth", code))
	return err
}
