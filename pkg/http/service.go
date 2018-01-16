package http

import (
	"code.ysitd.cloud/component/account/pkg/http/handler/pages"
	"code.ysitd.cloud/gin/utils/interfaces"
	"github.com/gin-gonic/gin"
)

const (
	endpointLoginForm   = "login_form"
	endpointLoginSubmit = "login_submit"
	endpointUserInfo    = "user_info"
)

func (s *service) init() {
	s.collector.RegisterHttp(endpointLoginForm, []string{})
	s.collector.RegisterHttp(endpointLoginSubmit, []string{"user"})
	s.collector.RegisterHttp(endpointUserInfo, []string{"user"})
}

func (s *service) CreateService() (app interfaces.Engine) {
	app = gin.Default()
	{
		s.bindMiddleware(app)
		s.registerLoginRoute(app)
		pages.Register(app)
		s.registerAPI(app)
		s.registerOAuth(app)
	}
	register(app)
	return
}
