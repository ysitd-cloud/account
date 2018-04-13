package user

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	AvatarUrl   string `json:"avatar_url"`
	password    []byte
}

func (u *User) ValidatePassword(password string) bool {
	return bcrypt.CompareHashAndPassword(u.password, []byte(password)) == nil
}

func (u *User) ChangePassword(password string) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	u.password = hash

	return
}
