package oauth

import (
	"log"

	"github.com/RangelReale/osin"
	"github.com/gin-gonic/gin"
	"github.com/tonyhhyip/go-di-container"
	"github.com/ysitd-cloud/account/http/middlewares"
)

func HandleAuthorize(c *gin.Context) {
	kernel := c.MustGet("kernel").(container.Kernel)
	server := kernel.Make("osin.server").(*osin.Server)
	defer server.Storage.Close()

	resp := server.NewResponse()
	defer resp.Close()

	if ar := server.HandleAuthorizeRequest(resp, c.Request); ar != nil {

		c.Set("osin.request", ar)
		c.Set("osin.response", resp)
		c.Next()
		return
	}
	if resp.IsError && resp.InternalError != nil {
		log.Printf("ERROR: %s\n", resp.InternalError)
	}
	osin.OutputJSON(resp, c.Writer, c.Request)
	c.Abort()
}

func HandleAuthorizeApprove(c *gin.Context) {
	kernel := c.MustGet("kernel").(container.Kernel)
	server := kernel.Make("osin.server").(*osin.Server)
	defer server.Storage.Close()

	req := c.MustGet("osin.request").(*osin.AuthorizeRequest)
	resp := c.MustGet("osin.response").(*osin.Response)
	session := middlewares.GetSession(c)

	req.Authorized = true
	req.UserData = session.Get("username")

	server.FinishAuthorizeRequest(resp, c.Request, req)
	osin.OutputJSON(resp, c.Writer, c.Request)
	c.Abort()
}
