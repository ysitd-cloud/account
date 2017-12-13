package oauth

import "golang.org/x/oauth2"

type AuthProvider interface {
	GetClientID() string
	GetClientSecret() string
	GetID() string
	GetName() string
	GetScopes() []string
	GetRedirectURL() string
	GetConfig() *oauth2.Config
	GetUserID(token *oauth2.Token) (string, error)
}

var providers = make(map[string]AuthProvider)

var booted = false

func RegisterProvider(id string, provider AuthProvider) {
	providers[id] = provider
}

func GetProvider(id string) AuthProvider {
	if !booted {
		boot()
	}
	provider, exists := providers[id]
	if exists {
		return provider
	}

	return nil
}
