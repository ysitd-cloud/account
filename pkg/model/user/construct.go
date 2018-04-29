package user

import (
	"database/sql"

	"golang.ysitd.cloud/db"
)

func ListFromDB(pool db.Opener) ([]*User, error) {
	conn, err := pool.Open()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := "SELECT username, display_name, email, avatar_uri FROM users"
	stmt, err := conn.Prepare(query)
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
		var username, displayName, email, avatarURL string
		if err := rows.Scan(&username, &displayName, &email, &avatarURL); err != nil {
			return nil, err
		}
		user := &User{
			Username:    username,
			DisplayName: displayName,
			Email:       email,
			AvatarUrl:   avatarURL,
			DB:          pool,
		}

		users = append(users, user)
	}

	return users, nil
}

func LoadFromDBWithUsername(pool db.Opener, username string) (*User, error) {
	conn, err := pool.Open()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	query := "SELECT username, display_name, email, avatar_url FROM users WHERE username = $1"
	row := conn.QueryRow(query, username)
	var displayName, email, avatarURL string
	if err := row.Scan(&username, &displayName, &email, &avatarURL); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	user := &User{
		Username:    username,
		DisplayName: displayName,
		Email:       email,
		AvatarUrl:   avatarURL,
		DB:          pool,
	}

	return user, nil
}
