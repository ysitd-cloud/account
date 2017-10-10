package oauth

import (
	"context"

	"strconv"

	githubClient "github.com/google/go-github/github"
	"github.com/ysitd-cloud/account/model"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type GithubAuthProvider struct {
	AbstractAuthProvider
}

func (provider *GithubAuthProvider) GetConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     provider.GetClientID(),
		ClientSecret: provider.GetClientSecret(),
		Endpoint:     github.Endpoint,
		RedirectURL:  provider.GetRedirectURL(),
		Scopes:       provider.GetScopes(),
	}
}

func (provider *GithubAuthProvider) GetUserID(token *oauth2.Token) (string, error) {
	config := provider.GetConfig()
	client := githubClient.NewClient(config.Client(context.Background(), token))
	user, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		return "", err
	}

	return strconv.Itoa(*user.ID), nil
}

func CreateGithubAuthProvider(provider *model.Provider) *GithubAuthProvider {
	return &GithubAuthProvider{
		AbstractAuthProvider: createAbstractProvider(provider),
	}
}
