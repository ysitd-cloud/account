package user

import (
	"context"
	"database/sql"

	"code.ysitd.cloud/common/go/db"
	"code.ysitd.cloud/component/account/pkg/modals"
)

type Store struct {
	DB db.Pool `inject:"db"`
}

func (s *Store) GetByUsername(ctx context.Context, username string) (u *User, err error) {
	conn, err := s.DB.Acquire()
	if err != nil {
		return
	}

	defer conn.Close()

	query := "SELECT username, display_name, email, avatar_url, password FROM users WHERE username = $1"
	row := conn.QueryRowContext(ctx, query, username)
	var displayName, email, avatarUrl, password string
	if err := row.Scan(&username, &displayName, &email, &avatarUrl, &password); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	user := &User{
		Username:    username,
		DisplayName: displayName,
		Email:       email,
		AvatarUrl:   avatarUrl,
		password:    []byte(password),
	}

	return user, nil
}

func (s *Store) ApplyChangePassword(ctx context.Context, u *User) (err error) {
	query := "UPDATE user_auth SET password = $1 WHERE username = $2"
	conn, err := s.DB.Acquire()
	if err != nil {
		return
	}

	defer conn.Close()

	result, err := conn.ExecContext(ctx, query, string(u.password), u.Username)

	if err != nil {
		return
	}

	if row, err := result.RowsAffected(); err != nil {
		return err
	} else if row != 1 {
		return &modals.IncorrectResultAffectedError{Row: row, Expected: 1}
	}

	return
}
