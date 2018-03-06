package proxy

import (
	"net/http"

	"code.ysitd.cloud/grpc/schema/account"
	"code.ysitd.cloud/grpc/schema/account/actions"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
)

func getUserInfo(c *gin.Context) {
	service := c.MustGet("service").(account.AccountServer)
	reply, err := service.GetUserInfo(c, &actions.GetUserInfoRequest{
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

	reply, err := service.ValidateUserPassword(c, &req)

	if err != nil {
		c.AbortWithError(http.StatusBadGateway, err)
		return
	}

	c.JSON(http.StatusOK, reply)
	c.Abort()
}
