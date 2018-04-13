package token

import (
	"context"
	"database/sql"
	"time"

	"code.ysitd.cloud/common/go/db"
	"code.ysitd.cloud/component/account/pkg/modals"
)

type Store struct {
	DB db.Pool `inject:"db"`
}

func (s *Store) GetToken(ctx context.Context, token string) (t *Token, err error) {
	conn, err := s.DB.Acquire()
	if err != nil {
		return
	}
	defer conn.Close()

	query := "SELECT access_token, extra, client, created_at, expires_in, scope FROM access FROM access_token = $1"
	row := conn.QueryRowContext(ctx, query, token)
	var instance Token
	var duration int64

	if err := row.Scan(
		&instance.AccessToken,
		&instance.Issuer,
		&instance.Audience,
		&instance.CreatedAt,
		&duration,
		&instance.Scope,
	); err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	instance.Expire = time.Duration(duration) * time.Second

	t = &instance

	return
}

func (s *Store) Revoke(ctx context.Context, token string) (err error) {
	conn, err := s.DB.Acquire()
	if err != nil {
		return
	}
	defer conn.Close()

	query := "DELETE FROM access WHERE access_token = $1"
	result, err := conn.ExecContext(ctx, query, token)
	if row, err := result.RowsAffected(); err != nil {
		return err
	} else if row != 1 {
		return &modals.IncorrectResultAffectedError{
			Expected: 1,
			Row:      row,
		}
	}
	return
}

func (s *Store) ExtendToken(ctx context.Context, token string, d time.Duration) (err error) {
	t, err := s.GetToken(ctx, token)
	if err != nil {
		return
	}

	conn, err := s.DB.Acquire()
	if err != nil {
		return
	}
	defer conn.Close()

	second := int(t.Expire / time.Second)

	query := `UPDATE access SET expires_in = $2 WHERE access_token = $1`

	result, err := conn.ExecContext(ctx, query, token, second)

	if row, err := result.RowsAffected(); err != nil {
		return err
	} else if row != 1 {
		return &modals.IncorrectResultAffectedError{
			Expected: 1,
			Row:      row,
		}
	}
	return
}
