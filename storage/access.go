package storage

import (
	"github.com/RangelReale/osin"
	"github.com/satori/go.uuid"
	"github.com/garyburd/redigo/redis"
)

func (s *Store) SaveAccess(data *osin.AccessData) (err error) {
	conn := s.Redis.Get()
	if err := conn.Err(); err != nil {
		return err
	}

	defer conn.Close()

	payload, err := encode(data)
	if err != nil {
		return err
	}

	accessID := uuid.NewV4().String()

	if _, err := conn.Do("SETEX", makeKey("access", accessID), data.ExpiresIn, string(payload)); err != nil {
		return err
	}

	if _, err := conn.Do("SETEX", makeKey("access_token", data.AccessToken), data.ExpiresIn, accessID); err != nil {
		return err
	}

	_, err = conn.Do("SETEX", makeKey("refresh_token", data.RefreshToken), data.ExpiresIn, accessID)
	return err
}

// LoadAccess gets access data with given access token
func (s *Store) LoadAccess(token string) (*osin.AccessData, error) {
	return s.loadAccessByKey(makeKey("access_token", token))
}

// RemoveAccess deletes AccessData with given access token
func (s *Store) RemoveAccess(token string) error {
	return s.removeAccessByKey(makeKey("access_token", token))
}

// LoadRefresh gets access data with given refresh token
func (s *Store) LoadRefresh(token string) (*osin.AccessData, error) {
	return s.loadAccessByKey(makeKey("refresh_token", token))
}

// RemoveRefresh deletes AccessData with given refresh token
func (s *Store) RemoveRefresh(token string) error {
	return s.removeAccessByKey(makeKey("refresh_token", token))
}

func (s *Store) removeAccessByKey(key string) error {
	conn := s.Redis.Get()
	if err := conn.Err(); err != nil {
		return err
	}

	defer conn.Close()

	accessID, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return err
	}

	access, err := s.loadAccessByKey(key)
	if err != nil {
		return err
	}

	if access == nil {
		return nil
	}

	accessKey := makeKey("access", accessID)
	if _, err := conn.Do("DEL", accessKey); err != nil {
		return err
	}

	accessTokenKey := makeKey("access_token", access.AccessToken)
	if _, err := conn.Do("DEL", accessTokenKey); err != nil {
		return err
	}

	refreshTokenKey := makeKey("refresh_token", access.RefreshToken)
	_, err = conn.Do("DEL", refreshTokenKey)
	return err
}

func (s *Store) loadAccessByKey(key string) (*osin.AccessData, error) {
	conn := s.Redis.Get()
	if err := conn.Err(); err != nil {
		return nil, err
	}

	defer conn.Close()

	var (
		rawAuthGob interface{}
		err        error
	)

	if rawAuthGob, err = conn.Do("GET", key); err != nil {
		return nil, err
	}
	if rawAuthGob == nil {
		return nil, nil
	}

	accessID, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	accessIDKey := makeKey("access", accessID)
	accessGob, err := redis.Bytes(conn.Do("GET", accessIDKey))
	if err != nil {
		return nil, err
	}

	var access osin.AccessData
	if err := decode(accessGob, &access); err != nil {
		return nil, err
	}

	ttl, err := redis.Int(conn.Do("TTL", accessIDKey))
	if err != nil {
		return nil, err
	}

	access.ExpiresIn = int32(ttl)

	access.Client, err = s.GetClient(access.Client.GetId())
	if err != nil {
		return nil, err
	}

	if access.AuthorizeData != nil && access.AuthorizeData.Client != nil {
		access.AuthorizeData.Client, err = s.GetClient(access.AuthorizeData.Client.GetId())
		if err != nil {
			return nil, err
		}
	}

	return &access, nil
}
