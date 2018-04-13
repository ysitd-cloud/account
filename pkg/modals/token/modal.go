package token

import "time"

type Token struct {
	AccessToken string        `json:"access_token"`
	Issuer      string        `json:"issuer"`
	Audience    string        `json:"audience"`
	CreatedAt   time.Time     `json:"created_at"`
	Expire      time.Duration `json:"expire"`
	Scope       string        `json:"scope"`
}

// ExpireAt returns the expiration date
func (t *Token) ExpireAt() time.Time {
	return t.CreatedAt.Add(time.Duration(t.Expire) * time.Second)
}

// IsExpiredAt returns true if access expires at time 't'
func (t *Token) IsExpiredAt(tm time.Time) bool {
	return t.ExpireAt().Before(tm)
}

// IsExpired returns true if access expired
func (t *Token) IsExpired() bool {
	return t.IsExpiredAt(time.Now())
}

// Extend extend token duration
func (t *Token) Extend(d time.Duration) {
	t.Expire += d
}
