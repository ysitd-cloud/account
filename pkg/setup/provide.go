package setup

import (
	"net/http"

	"code.ysitd.cloud/component/account/pkg/server/public"
)

var publicService public.Provider

func GetPublicServiceHandler() http.Handler {
	return publicService.CreateHandler()
}
