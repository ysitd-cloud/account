package handler

import (
	"github.com/RangelReale/osin"
	"gopkg.in/gin-gonic/gin.v1"
	"fmt"
)

func HandleTokenRequest(c *gin.Context) {
	server := c.MustGet("osin.server").(*osin.Server)
	resp := server.NewResponse()
	defer resp.Close()

	if req := server.HandleAccessRequest(resp, c.Request); req != nil {
		switch req.Type {
		case osin.AUTHORIZATION_CODE:
			req.Authorized = true
		case osin.REFRESH_TOKEN:
			req.Authorized = true
		case osin.PASSWORD:
			if req.Username == "test" && req.Password == "test" {
				req.Authorized = true
			}
		case osin.CLIENT_CREDENTIALS:
			req.Authorized = true
		}
		resp.Output["user"] = req.AuthorizeData.UserData.(string)
		req.UserData = resp.Output["user"]
		server.FinishAccessRequest(resp, c.Request, req)
	}

	if resp.IsError && resp.InternalError != nil {
		fmt.Printf("ERROR: %s\n", resp.InternalError)
	}
	osin.OutputJSON(resp, c.Writer, c.Request)
}
