package user

import "golang.ysitd.cloud/db"

type User struct {
	Username    string    `json:"username"`
	DisplayName string    `json:"display_name"`
	Email       string    `json:"email"`
	AvatarUrl   string    `json:"avatar_url"`
	DB          db.Opener `json:"-"`
}
