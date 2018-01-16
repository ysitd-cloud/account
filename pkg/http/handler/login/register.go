package login

import (
	"code.ysitd.cloud/component/account/pkg/metrics"
	"github.com/gin-gonic/gin"
)

const (
	EndpointLoginForm   = "login_form"
	EndpointLoginSubmit = "login_submit"
)

func Register(app gin.IRoutes, collector metrics.Collector) {
	app.GET("/login", basicForm(collector))
	app.POST("/login", basicSubmit(collector))
}
