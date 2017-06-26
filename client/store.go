package client

import (
	"gopkg.in/oauth2.v3"
	"github.com/jinzhu/gorm"
)

type OAuthClientStore struct {
	Database *gorm.DB
}

func NewOAuthClientStore(db *gorm.DB) OAuthClientStore {
	return OAuthClientStore{
		db,
	}
}

func (store OAuthClientStore) GetByID(id string) (oauth2.ClientInfo, error) {
	client := Client{}
	store.Database.Where("client_id = ?", id).First(&client)
	return client, nil
}
