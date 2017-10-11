package oauth

import (
	"database/sql"

	"github.com/ysitd-cloud/account/model"
	container "github.com/ysitd-cloud/account/providers"
)

func init() {
	db := container.Kernel.Make("db").(*sql.DB)

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
