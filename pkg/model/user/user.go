package user

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func (user *User) ValidatePassword(password string) bool {
	var hash string
	query := "SELECT password FROM user_auth WHERE username = $1"

	db := user.getDB()
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

	db := user.getDB()
	defer db.Close()

	_, err = db.Exec(query, string(hash), user.Username)
	return err
}

func (user *User) getDB() *sql.DB {
	db, _ := user.DB.Acquire()
	return db
}
