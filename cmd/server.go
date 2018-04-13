package main

import (
	"net/http"
	"os"

	"code.ysitd.cloud/component/account/pkg/setup"
)

func main() {
	{
		handler := setup.GetPublicServiceHandler()
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		http.ListenAndServe(":"+port, handler)
	}
}
