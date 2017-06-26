package client

import "github.com/jinzhu/gorm"

type Client struct {
	gorm.Model
	ClientID string `gorm:"primary_key;type=UUID;not null"`
	Secret string `gorm:"not null"`
	Domain string `gorm:"not null;index"`
	UserID string `gorm:"not null;index"`
}

func (client Client) GetID() string {
	return client.ClientID
}

func (client Client) GetSecret() string {
	return client.Secret
}

func (client Client) GetDomain() string {
	return client.Domain
}

func (client Client) GetUserID() string {
	return client.UserID
}
