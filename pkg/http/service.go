package http

import (
	"code.ysitd.cloud/component/account/pkg/http/handler/login"
	"code.ysitd.cloud/gin/utils/interfaces"
	"github.com/gin-gonic/gin"
)

func (s *service) init() {
	s.collector.RegisterHttp(login.EndpointLoginForm, []string{})
	s.collector.RegisterHttp(login.EndpointLoginSubmit, []string{"user"})
}

func (s *service) CreateService() (app interfaces.Engine) {
	app = gin.Default()
	Register(app, s.collector)
	return
}
