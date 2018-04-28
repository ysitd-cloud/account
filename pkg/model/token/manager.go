package token

import (
	"database/sql"
	"time"
)

func (m *manager) GetToken(token string) (*Token, error) {
	db, err := m.pool.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT access_token, extra, client, created_at, expires_in, scope FROM access FROM access_token = $1"
	row := db.QueryRow(query, token)
	var t Token
	var duration int64

	if err := row.Scan(&t.AccessToken, &t.Issuer, &t.Audience, &t.CreatedAt, &duration, &t.Scope); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	t.Expire = time.Duration(duration) * time.Second

	return &t, nil
}

func (m *manager) Revoke(token string) error {
	db, err := m.pool.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "DELETE FROM access WHERE access_token = $1"
	_, err = db.Exec(query, token)
	return err
}

func (m *manager) ExtendToken(token string, duration time.Duration) error {
	t, err := m.GetToken(token)
	if err != nil {
		return err
	}
	t.Expire += duration

	return m.updateToken(t)
}

func (m *manager) updateToken(t *Token) error {
	db, err := m.pool.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	second := int(t.Expire / time.Second)

	query := `UPDATE access SET (extra, client, created_at, expires_in, scope) = ($2, $3, $4, $5, $6) WHERE access_token = $1`
	_, err = db.Exec(query, t.AccessToken, t.Issuer, t.Audience, t.CreatedAt, second, t.Scope)
	return err
}
