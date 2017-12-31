package proxy

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/grpc-schema/account"
	"github.com/ysitd-cloud/grpc-schema/account/actions"
	"golang.org/x/net/context"
)

func getUserInfo(c *gin.Context) {
	service := c.MustGet("service").(account.AccountServer)
	reply, err := service.GetUserInfo(context.Background(), &actions.GetUserInfoRequest{
		Username: c.Param("username"),
	})

	if err != nil {
		c.AbortWithError(http.StatusBadGateway, err)
		return
	}

	c.JSON(http.StatusOK, reply)
	c.Abort()
}

func validateUserPassword(c *gin.Context) {
	service := c.MustGet("service").(account.AccountServer)

	var req actions.ValidateUserRequest
	c.BindJSON(&req)

	reply, err := service.ValidateUserPassword(context.Background(), &req)

	if err != nil {
		c.AbortWithError(http.StatusBadGateway, err)
		return
	}

	c.JSON(http.StatusOK, reply)
	c.Abort()
}
