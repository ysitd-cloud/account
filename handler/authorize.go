package handler

import (
	"github.com/RangelReale/osin"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"github.com/ysitd-cloud/account/middlewares"
)

func HandleAuthorize(c *gin.Context) {
	server := c.MustGet("osin.server").(*osin.Server)
	log.Println("Middleware:HandleAuthorize")
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
	log.Println("Middleware:HandleAuthorizeApprove")

	server := c.MustGet("osin.server").(*osin.Server)
	req := c.MustGet("osin.request").(*osin.AuthorizeRequest)
	resp := c.MustGet("osin.response").(*osin.Response)
	session := middlewares.GetSession(c)

	req.Authorized = true
	req.UserData = session.Get("username")

	server.FinishAuthorizeRequest(resp, c.Request, req)
	osin.OutputJSON(resp, c.Writer, c.Request)
	c.Abort()
}
