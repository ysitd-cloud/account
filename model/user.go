package model

import (
	"database/sql"
	"github.com/ysitd-cloud/account/setup"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	AvatarUrl   string `json:"avatar_url"`
}

func LoadUserFromDBWithUsername(db *sql.DB, username string) (*User, error) {
	query := "SELECT username, display_name, email, avatar_url FROM users WHERE username = $1"
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(username)
	return LoadFromRow(row)
}

func LoadFromRow(row *sql.Row) (*User, error) {
	var username, displayName, email, avatarUrl string
	if err := row.Scan(&username, &displayName, &email, &avatarUrl); err != nil {
		return nil, err
	}
	user := &User{
		Username:    username,
		DisplayName: displayName,
		Email:       email,
		AvatarUrl:   avatarUrl,
	}

	return user, nil
}

func (user *User) ValidatePassword(password string) bool {
	var hash string
	query := "SELECT password FROM user_auth WHERE username = $1"
	db, err := setup.SetupDB()
	if err != nil {
		return false
	}

	row := db.QueryRow(query, user.Username)
	if err := row.Scan(&hash); err != nil {
		return false
	}

	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
