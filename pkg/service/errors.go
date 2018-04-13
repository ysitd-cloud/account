package service

import "errors"

var (
	ErrIncorrectPassword = errors.New("incorrect password")
	ErrUserNotExists     = errors.New("user not existss")
)
