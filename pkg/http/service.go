package http

import (
	"sync"

	"code.ysitd.cloud/auth/account/pkg/http/handler/pages"
	"code.ysitd.cloud/gin/utils/interfaces"
	"github.com/gin-gonic/gin"
)

const (
	endpointLoginForm   = "login_form"
	endpointLoginSubmit = "login_submit"
	endpointUserInfo    = "user_info"
)

var once sync.Once

func (s *service) init() {
	once.Do(func() {
		s.Collector.RegisterHTTP(endpointLoginForm, []string{})
		s.Collector.RegisterHTTP(endpointLoginSubmit, []string{"user"})
		s.Collector.RegisterHTTP(endpointUserInfo, []string{"user"})
	})
}

func (s *service) CreateService() (app interfaces.Engine) {
	app = gin.Default()
	{
		app.Use(func(c *gin.Context) {
			c.Set("kernel", s.app)
			c.Next()
		})
		s.bindMiddleware(app)
		s.LoginHandler.RegisterRoutes(app)
		pages.Register(app)
		s.registerAPI(app)
		s.registerOAuth(app)
	}
	register(app)
	return
}
