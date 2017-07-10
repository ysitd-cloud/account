package middlewares

import (
	"github.com/RangelReale/osin"
	"gopkg.in/gin-gonic/gin.v1"
)

func HandleAuthorize(server *osin.Server) (handlerFunc gin.HandlerFunc) {
	return func (c *gin.Context) {
		resp := server.NewResponse()
		defer resp.Close()

		if ar := server.HandleAuthorizeRequest(resp, c.Request); ar != nil {
			c.Next()
			return
		}
		if resp.IsError && resp.InternalError != nil {
			c.AbortWithError(500, resp.InternalError)
		}
		osin.OutputJSON(resp, c.Writer, c.Request)
		c.Abort()
	}
}
