package proxy

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/grpc-schema/account"
	"github.com/ysitd-cloud/grpc-schema/account/actions"
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
