package user

import (
	"database/sql"

	"github.com/ysitd-cloud/account/pkg/utils"
)

func ListFromDB(pool utils.DatabasePool) ([]*User, error) {
	db, err := pool.Acquire()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT username, display_name, email, avatar_uri FROM users"
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var users []*User

	defer rows.Close()
	for rows.Next() {
		var username, displayName, email, avatarUrl string
		if err := rows.Scan(&username, &displayName, &email, &avatarUrl); err != nil {
			return nil, err
		}
		user := &User{
			Username:    username,
			DisplayName: displayName,
			Email:       email,
			AvatarUrl:   avatarUrl,
			DB:          pool,
		}

		users = append(users, user)
	}

	return users, nil
}

func LoadFromDBWithUsername(pool utils.DatabasePool, username string) (*User, error) {
	db, err := pool.Acquire()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT username, display_name, email, avatar_url FROM users WHERE username = $1"
	row := db.QueryRow(query, username)
	var displayName, email, avatarUrl string
	if err := row.Scan(&username, &displayName, &email, &avatarUrl); err != nil {
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
		DB:          pool,
	}

	return user, nil
}
