package storage

import (
	"database/sql"
	"errors"

	"github.com/RangelReale/osin"
	"github.com/garyburd/redigo/redis"
)

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
		return err
	}

	db, err := s.DB.Acquire()
	if err != nil {
		return err
	}

	defer db.Close()

	tx, err := db.Begin()
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
			return errors.New(rbe.Error())
		}
		return errors.New(err.Error())
	}
	return nil
}

// LoadAccess gets access data with given access token
func (s *Store) LoadAccess(code string) (*osin.AccessData, error) {
	var extra, cid, prevAccessToken, authorizeCode string
	var result osin.AccessData

	db, err := s.DB.Acquire()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	query := "SELECT client, authorize, previous, access_token, refresh_token, expires_in, scope, redirect_uri, created_at, extra FROM access WHERE access_token=$1 LIMIT 1"
	if err := db.QueryRow(
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
		return nil, errNotFound
	} else if err != nil {
		return nil, errors.New(err.Error())
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
func (s *Store) RemoveAccess(code string) error {
	db, err := s.DB.Acquire()
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("DELETE FROM refresh WHERE token=$1", code)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// LoadRefresh gets access data with given refresh token
func (s *Store) LoadRefresh(code string) (*osin.AccessData, error) {
	db, err := s.DB.Acquire()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	row := db.QueryRow("SELECT access FROM refresh WHERE token=$1 LIMIT 1", code)
	var access string
	if err := row.Scan(&access); err == sql.ErrNoRows {
		return nil, errNotFound
	} else if err != nil {
		return nil, errors.New(err.Error())
	}
	return s.LoadAccess(access)
}

// RemoveRefresh deletes AccessData with given refresh token
func (s *Store) RemoveRefresh(token string) error {
	db, err := s.DB.Acquire()
	if err != nil {
		return err
	}

	defer db.Close()

	_, err = db.Exec("DELETE FROM refresh WHERE token=$1", token)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
