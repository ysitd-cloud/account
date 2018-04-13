package setup

import (
	"net/http"
	"time"

	"code.ysitd.cloud/component/account/pkg/config"
	"github.com/facebookgo/inject"
)

func setUpHttpClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}

func injectHttpClient(_ *config.Config, graph *inject.Graph) {
	graph.Provide(&inject.Object{
		Value: setUpHttpClient(),
	})
}
