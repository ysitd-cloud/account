package storage

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/RangelReale/osin"
)

var errNotFound = errors.New("not found")

func (s *Store) GetClient(id string) (osin.Client, error) {
	db, err := s.DB.Acquire()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	row := db.QueryRow("SELECT id, secret, redirect_uri, extra FROM client WHERE id=$1", id)
	var c osin.DefaultClient
	var extra string

	if err := row.Scan(&c.Id, &c.Secret, &c.RedirectUri, &extra); err == sql.ErrNoRows {
		return nil, errNotFound
	} else if err != nil {
		return nil, err
	}
	c.UserData = extra
	return &c, nil
}

// UpdateClient updates the client (identified by it's id) and replaces the values with the values of client.
func (s *Store) UpdateClient(c osin.Client) error {
	data, err := assertToString(c.GetUserData())
	if err != nil {
		return err
	}

	db, err := s.DB.Acquire()
	if err != nil {
		return err
	}

	defer db.Close()

	query := "UPDATE client SET (secret, redirect_uri, extra) = ($2, $3, $4) WHERE id = $1"
	if _, err := db.Exec(query, c.GetId(), c.GetSecret(), c.GetRedirectUri(), data); err != nil {
		return err
	}
	return nil
}

// CreateClient stores the client in the database and returns an error, if something went wrong.
func (s *Store) CreateClient(c osin.Client) error {
	data, err := assertToString(c.GetUserData())
	if err != nil {
		return err
	}

	db, err := s.DB.Acquire()
	if err != nil {
		return err
	}

	defer db.Close()

	query := "INSERT INTO client (id, secret, redirect_uri, extra) VALUES ($1, $2, $3, $4)"
	if _, err := db.Exec(query, c.GetId(), c.GetSecret(), c.GetRedirectUri(), data); err != nil {
		return err
	}
	return nil
}

// RemoveClient removes a client (identified by id) from the database. Returns an error if something went wrong.
func (s *Store) RemoveClient(id string) (err error) {
	db, err := s.DB.Acquire()
	if err != nil {
		return err
	}

	defer db.Close()
	if _, err = db.Exec("DELETE FROM client WHERE id=$1", id); err != nil {
		return err
	}
	return nil
}

func assertToString(in interface{}) (string, error) {
	var ok bool
	var data string
	if in == nil {
		return "", nil
	} else if data, ok = in.(string); ok {
		return data, nil
	} else if str, ok := in.(fmt.Stringer); ok {
		return str.String(), nil
	}
	return "", errors.New("could not assert to string")
}
