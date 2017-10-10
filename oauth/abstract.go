package oauth

import "github.com/ysitd-cloud/account/model"

type AbstractAuthProvider struct {
	Provider *model.Provider
}

func (a *AbstractAuthProvider) GetClientID() string {
	return a.Provider.ClientID
}

func (a *AbstractAuthProvider) GetClientSecret() string {
	return a.Provider.ClientSecret
}

func (a *AbstractAuthProvider) GetID() string {
	return a.Provider.ID
}

func (a *AbstractAuthProvider) GetName() string {
	return a.Provider.Name
}

func (a *AbstractAuthProvider) GetScopes() []string {
	return a.Provider.Scopes
}

func (a *AbstractAuthProvider) GetRedirectURL() string {
	return a.Provider.RedirectURL
}

func createAbstractProvider(provider *model.Provider) AbstractAuthProvider {
	return AbstractAuthProvider{
		Provider: provider,
	}
}
