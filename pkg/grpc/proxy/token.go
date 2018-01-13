package proxy

import (
	"net/http"

	"code.ysitd.cloud/grpc/schema/account"
	"code.ysitd.cloud/grpc/schema/account/actions"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func getTokenInfo(c *gin.Context) {
	service := c.MustGet("service").(account.AccountServer)
	reply, err := service.GetTokenInfo(context.Background(), &actions.GetTokenInfoRequest{
		Token: c.Param("token"),
	})

	if err != nil {
		c.AbortWithError(http.StatusBadGateway, err)
		return
	}

	c.JSON(http.StatusOK, reply)
	c.Abort()
}
