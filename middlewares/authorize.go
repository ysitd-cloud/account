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
			c.Set("osin.request", ar)
			c.Set("osin.response", resp)
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

func HandleAuthorizeApprove(server *osin.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.MustGet("osin.request").(*osin.AuthorizeRequest)
		resp := c.MustGet("osin.response").(*osin.Response)
		req.Authorized = true
		server.FinishAuthorizeRequest(resp, c.Request, req)
	}
}
