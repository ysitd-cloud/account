package session

import (
	"code.ysitd.cloud/component/account/pkg/modals/user"
	"github.com/dgrijalva/jwt-go"
)

type Session struct {
	jwt.StandardClaims
	User *user.User `json:"user"`
}
