package oauth

import (
	"fmt"

	"github.com/RangelReale/osin"
	"github.com/gin-gonic/gin"
	"github.com/tonyhhyip/go-di-container"
)

func HandleTokenRequest(c *gin.Context) {
	kernel := c.MustGet("kernel").(container.Kernel)
	server := kernel.Make("osin.server").(*osin.Server)
	defer server.Storage.Close()

	resp := server.NewResponse()
	defer resp.Close()

	if req := server.HandleAccessRequest(resp, c.Request); req != nil {
		switch req.Type {
		case osin.AUTHORIZATION_CODE:
			req.Authorized = true
			resp.Output["user"] = req.AuthorizeData.UserData.(string)
			req.UserData = resp.Output["user"]
		case osin.REFRESH_TOKEN:
			req.Authorized = true
		case osin.PASSWORD:
			if req.Username == "test" && req.Password == "test" {
				req.Authorized = true
			}
		case osin.CLIENT_CREDENTIALS:
			req.Authorized = true
		}
		server.FinishAccessRequest(resp, c.Request, req)
	}

	if resp.IsError && resp.InternalError != nil {
		fmt.Printf("ERROR: %s\n", resp.InternalError)
	}
	osin.OutputJSON(resp, c.Writer, c.Request)
}
