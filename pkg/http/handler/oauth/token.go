package oauth

import (
	"fmt"

	"code.ysitd.cloud/auth/account/third_party/forked/github.com/RangelReale/osin"
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
		case osin.AuthorizationCode:
			req.Authorized = true
			resp.Output["user"] = req.AuthorizeData.UserData.(string)
			req.UserData = resp.Output["user"]
		case osin.RefreshToken:
			req.Authorized = true
		}
		server.FinishAccessRequest(resp, c.Request, req)
	}

	if resp.IsError && resp.InternalError != nil {
		fmt.Printf("ERROR: %s\n", resp.InternalError)
	}
	osin.OutputJSON(resp, c.Writer, c.Request)
}
