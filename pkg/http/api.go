package http

import (
	"net/http"

	"code.ysitd.cloud/common/go/db"
	"code.ysitd.cloud/component/account/pkg/http/middlewares"
	"code.ysitd.cloud/component/account/pkg/model/user"
	"github.com/RangelReale/osin"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

func (s *service) registerAPI(app gin.IRouter) {
	group := app.Group("/api")
	v1 := group.Group("v1", middlewares.ContainsAuthHeader, middlewares.BearerToken)

	{
		v1.GET("/user/info", func(c *gin.Context) {
			access := c.MustGet("oauth.access").(*osin.AccessData)
			approved := access.UserData.(string)
			labels := prometheus.Labels{
				"user": approved,
			}
			finish, err := s.collector.InvokeHTTP(endpointUserInfo, labels)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}

			defer func() {
				finish <- c.Writer.Status()
				close(finish)
			}()

			pool := s.app.Make("db.pool").(db.Pool)

			instance, err := user.LoadFromDBWithUsername(pool, approved)
			if err != nil {
				c.AbortWithError(http.StatusNotFound, err)
			} else {
				c.JSON(http.StatusOK, instance)
				c.Abort()
			}
		})
	}
}
