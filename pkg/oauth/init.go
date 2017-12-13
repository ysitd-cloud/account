package oauth

import (
	"database/sql"

	"github.com/ysitd-cloud/account/pkg/kernel"
	"github.com/ysitd-cloud/account/pkg/model"
)

func boot() {
	booted = true
	db := kernel.Kernel.Make("db").(*sql.DB)

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
