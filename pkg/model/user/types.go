package user

import "code.ysitd.cloud/component/account/pkg/utils"

type User struct {
	Username    string             `json:"username"`
	DisplayName string             `json:"display_name"`
	Email       string             `json:"email"`
	AvatarUrl   string             `json:"avatar_url"`
	DB          utils.DatabasePool `json:"-"`
}
