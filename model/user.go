package model

import (
	"database/sql"

	"github.com/ysitd-cloud/account/providers"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	AvatarUrl   string `json:"avatar_url"`
}

func ListUserFromDB(db *sql.DB) ([]*User, error) {
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
		}

		users = append(users, user)
	}

	return users, nil
}

func LoadUserFromDBWithUsername(db *sql.DB, username string) (*User, error) {
	query := "SELECT username, display_name, email, avatar_url FROM users WHERE username = $1"
	row := db.QueryRow(query, username)
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

	db := providers.Kernel.Make("db").(*sql.DB)
	defer db.Close()

	row := db.QueryRow(query, user.Username)
	if err := row.Scan(&hash); err != nil {
		return false
	}

	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func (user *User) ChangePassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := "UPDATE user_auth SET password = $1 WHERE username = $2"

	db := providers.Kernel.Make("db").(*sql.DB)
	defer db.Close()

	_, err = db.Exec(query, string(hash), user.Username)
	return err
}
