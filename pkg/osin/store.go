package osin

import (
	"database/sql"
	"errors"

	"code.ysitd.cloud/common/go/db"
	"github.com/RangelReale/osin"
	"github.com/garyburd/redigo/redis"
)

type Store struct {
	DB    db.Pool     `inject:"db"`
	Redis *redis.Pool `inject:""`
}

func (s *Store) Clone() osin.Storage {
	return s
}

func (s *Store) Close() {}

func (s *Store) GetClient(id string) (client osin.Client, err error) {
	conn, err := s.DB.Acquire()
	if err != nil {
		return
	}

	defer conn.Close()
	row := conn.QueryRow("SELECT id, secret, redirect_uri, extra FROM client WHERE id=$1", id)
	var c osin.DefaultClient
	var extra string

	if err := row.Scan(&c.Id, &c.Secret, &c.RedirectUri, &extra); err == sql.ErrNoRows {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}
	c.UserData = extra

	client = &c

	return
}

// UpdateClient updates the client (identified by it's id) and replaces the values with the values of client.
func (s *Store) UpdateClient(c osin.Client) (err error) {
	data, err := assertToString(c.GetUserData())
	if err != nil {
		return
	}

	conn, err := s.DB.Acquire()
	if err != nil {
		return
	}

	defer conn.Close()

	query := "UPDATE client SET (secret, redirect_uri, extra) = ($2, $3, $4) WHERE id = $1"
	if _, err := conn.Exec(query, c.GetId(), c.GetSecret(), c.GetRedirectUri(), data); err != nil {
		return err
	}
	return
}

// CreateClient stores the client in the database and returns an error, if something went wrong.
func (s *Store) CreateClient(c osin.Client) (err error) {
	data, err := assertToString(c.GetUserData())
	if err != nil {
		return
	}

	conn, err := s.DB.Acquire()
	if err != nil {
		return
	}

	defer conn.Close()

	query := "INSERT INTO client (id, secret, redirect_uri, extra) VALUES ($1, $2, $3, $4)"
	if _, err := conn.Exec(query, c.GetId(), c.GetSecret(), c.GetRedirectUri(), data); err != nil {
		return err
	}
	return
}

// RemoveClient removes a client (identified by id) from the database. Returns an error if something went wrong.
func (s *Store) RemoveClient(id string) (err error) {
	conn, err := s.DB.Acquire()
	if err != nil {
		return
	}

	defer conn.Close()
	if _, err = conn.Exec("DELETE FROM client WHERE id=$1", id); err != nil {
		return
	}
	return nil
}

func (s *Store) SaveAccess(data *osin.AccessData) (err error) {
	prev := ""
	authorizeData := &osin.AuthorizeData{}

	if data.AccessData != nil {
		prev = data.AccessData.AccessToken
	}

	if data.AuthorizeData != nil {
		authorizeData = data.AuthorizeData
	}

	extra, err := assertToString(data.UserData)
	if err != nil {
		return
	}

	conn, err := s.DB.Acquire()
	if err != nil {
		return
	}

	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		return errors.New(err.Error())
	}

	if data.RefreshToken != "" {
		if err := s.saveRefresh(tx, data.RefreshToken, data.AccessToken); err != nil {
			return err
		}
	}

	if data.Client == nil {
		return errors.New("data.Client must not be nil")
	}

	query := "INSERT INTO access (client, authorize, previous, access_token, refresh_token, expires_in, scope, redirect_uri, created_at, extra) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
	_, err = tx.Exec(query, data.Client.GetId(), authorizeData.Code, prev, data.AccessToken, data.RefreshToken, data.ExpiresIn, data.Scope, data.RedirectUri, data.CreatedAt, extra)
	if err != nil {
		if rbe := tx.Rollback(); rbe != nil {
			return errors.New(rbe.Error())
		}
		return errors.New(err.Error())
	}

	if err = tx.Commit(); err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (s *Store) saveRefresh(tx *sql.Tx, refresh, access string) (err error) {
	query := "INSERT INTO refresh (token, access) VALUES ($1, $2)"
	_, err = tx.Exec(query, refresh, access)
	if err != nil {
		if rbe := tx.Rollback(); rbe != nil {
			return rbe
		}
		return err
	}
	return
}

// LoadAccess gets access data with given access token
func (s *Store) LoadAccess(code string) (*osin.AccessData, error) {
	var extra, cid, prevAccessToken, authorizeCode string
	var result osin.AccessData

	conn, err := s.DB.Acquire()
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	query := "SELECT client, authorize, previous, access_token, refresh_token, expires_in, scope, redirect_uri, created_at, extra FROM access WHERE access_token=$1 LIMIT 1"
	if err := conn.QueryRow(
		query,
		code,
	).Scan(
		&cid,
		&authorizeCode,
		&prevAccessToken,
		&result.AccessToken,
		&result.RefreshToken,
		&result.ExpiresIn,
		&result.Scope,
		&result.RedirectUri,
		&result.CreatedAt,
		&extra,
	); err == sql.ErrNoRows {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	result.UserData = extra
	client, err := s.GetClient(cid)
	if err != nil {
		return nil, err
	}

	result.Client = client
	result.AuthorizeData, _ = s.LoadAuthorize(authorizeCode)
	prevAccess, _ := s.LoadAccess(prevAccessToken)
	result.AccessData = prevAccess
	return &result, nil
}

// RemoveAccess deletes AccessData with given access token
func (s *Store) RemoveAccess(code string) (err error) {
	conn, err := s.DB.Acquire()
	if err != nil {
		return
	}

	defer conn.Close()

	_, err = conn.Exec("DELETE FROM refresh WHERE token=$1", code)
	if err != nil {
		return
	}
	return
}

// LoadRefresh gets access data with given refresh token
func (s *Store) LoadRefresh(code string) (ad *osin.AccessData, err error) {
	conn, err := s.DB.Acquire()
	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow("SELECT access FROM refresh WHERE token=$1 LIMIT 1", code)
	var access string
	if err := row.Scan(&access); err == sql.ErrNoRows {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}
	return s.LoadAccess(access)
}

// RemoveRefresh deletes AccessData with given refresh token
func (s *Store) RemoveRefresh(token string) (err error) {
	conn, err := s.DB.Acquire()
	if err != nil {
		return
	}

	defer conn.Close()

	_, err = conn.Exec("DELETE FROM refresh WHERE token=$1", token)
	if err != nil {
		return
	}
	return
}

func (s *Store) SaveAuthorize(data *osin.AuthorizeData) (err error) {
	conn := s.Redis.Get()
	if rerr := conn.Err(); rerr != nil {
		return rerr
	}

	defer conn.Close()

	payload, err := encode(data)
	if err != nil {
		return
	}

	_, err = conn.Do("SETEX", makeKey("auth", data.Code), data.ExpiresIn, string(payload))
	return
}

func (s *Store) LoadAuthorize(code string) (ad *osin.AuthorizeData, err error) {
	conn := s.Redis.Get()
	if rerr := conn.Err(); rerr != nil {
		return nil, rerr
	}

	defer conn.Close()

	var (
		rawAuthGob interface{}
	)

	if rawAuthGob, err = conn.Do("GET", makeKey("auth", code)); err != nil {
		return
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
	if rerr := conn.Err(); rerr != nil {
		return rerr
	}

	defer conn.Close()

	_, err = conn.Do("DEL", makeKey("auth", code))
	return
}
