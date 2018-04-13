package public

import (
	"net/http"
	"os"

	"code.ysitd.cloud/component/account/pkg/server/public/session"
	"code.ysitd.cloud/component/account/pkg/service"
	"github.com/gorilla/handlers"
	"github.com/tonyhhyip/vodka"
)

type Provider struct {
	Service  *service.Service `inject:""`
	Manager  *session.Manager `inject:""`
	Renderer *Renderer        `inject:""`
}

func (p *Provider) CreateHandler() http.Handler {
	router := p.createRouter()
	handler := router.Handler()
	s := vodka.New("")
	s.StandBy(handler)
	return handlers.CombinedLoggingHandler(os.Stdout, s.Server.Handler)
}

func (p *Provider) createRouter() (r *vodka.Router) {
	r = vodka.NewRouter()
	r.Use(&securityMiddleware{
		FrameDeny:          true,
		BrowserXssFilter:   true,
		ContentTypeNosniff: true,
		STSSeconds:         5184000,
	})

	r.GET("/login", p.loginForm)
	r.POST("/login", p.loginHandle)
	return
}
