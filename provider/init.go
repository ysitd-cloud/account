package provider

import (
	"github.com/ysitd-cloud/account/model"
	"github.com/ysitd-cloud/account/setup"
)

func init() {
	db, err := setup.OpenDB()
	if err != nil {
		panic(err)
	}

	{
		provider, err := model.GetProviderByID(db, "github")
		if err != nil {
			panic(err)
		}

		if provider != nil {
			authProvider := CreateGithubAuthProvider(provider)
			RegisterProvider("github", authProvider)
		}
	}
}
