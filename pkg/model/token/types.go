package token

import (
	"time"

	"golang.ysitd.cloud/db"
)

type Token struct {
	AccessToken string        `json:"access_token"`
	Issuer      string        `json:"issuer"`
	Audience    string        `json:"audience"`
	CreatedAt   time.Time     `json:"created_at"`
	Expire      time.Duration `json:"expire"`
	Scope       string        `json:"scope"`
}

type Manager interface {
	GetToken(token string) (*Token, error)
	Revoke(token string) error
	ExtendToken(token string, duration time.Duration) error
}

type manager struct {
	pool db.Opener
}
