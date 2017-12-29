package token

import "time"

// IsExpired returns true if access expired
func (d *Token) IsExpired() bool {
	return d.IsExpiredAt(time.Now())
}

// IsExpiredAt returns true if access expires at time 't'
func (d *Token) IsExpiredAt(t time.Time) bool {
	return d.ExpireAt().Before(t)
}

// ExpireAt returns the expiration date
func (d *Token) ExpireAt() time.Time {
	return d.CreatedAt.Add(time.Duration(d.Expire) * time.Second)
}
