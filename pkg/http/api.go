package http

import (
	"net/http"

	"code.ysitd.cloud/auth/account/pkg/http/middlewares"
	"code.ysitd.cloud/auth/account/pkg/model/user"
	"code.ysitd.cloud/auth/account/third_party/forked/github.com/RangelReale/osin"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"golang.ysitd.cloud/db"
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
			done, err := s.Collector.InvokeHTTP(endpointUserInfo, labels)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}

			defer done(c.Writer.Status())

			pool := s.app.Make("db.pool").(db.Opener)

			instance, err := user.LoadFromDBWithUsername(c.Request.Context(), pool, approved)
			if err != nil {
				c.AbortWithError(http.StatusNotFound, err)
			} else {
				c.JSON(http.StatusOK, instance)
				c.Abort()
			}
		})
	}
}
