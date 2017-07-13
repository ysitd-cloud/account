package model

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	DisplayName string `json:"display_name"`
	Email string `json:"email"`
	AvatarUrl string `json:"avatar_url"`
	password string
	db *sql.DB
}

func LoadUserFromDBWithUsername(db *sql.DB, username string) (*User, error) {
	query := "SELECT username, password, display_name, email, avatar_url FROM users WHERE username = $1"
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(username)
	return LoadFromRow(row)
}

func LoadFromRow(row *sql.Row) (*User, error) {
	var username, displayName, email, avatarUrl, password string
	if err := row.Scan(&username, &password, &displayName, &email, &avatarUrl); err != nil {
		return nil, err
	}
	user := &User {
		Username: username,
		password: password,
		DisplayName: displayName,
		Email: email,
		AvatarUrl: avatarUrl,
	}

	return user, nil
}

func (user *User) ValidatePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.password), []byte(password)) == nil
}
