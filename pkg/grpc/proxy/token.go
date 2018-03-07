package proxy

import (
	"net/http"

	"code.ysitd.cloud/grpc/schema/account/actions"
	"github.com/tonyhhyip/vodka"
)

func (p *proxy) getTokenInfo(c *vodka.Context) {
	service := p.service
	token, err := vodka.String(c.UserValue("token"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"cause":   "fail to parse token from url",
			"message": err.Error(),
		})
		return
	}
	reply, err := service.GetTokenInfo(c, &actions.GetTokenInfoRequest{
		Token: token,
	})

	if err != nil {
		c.JSON(http.StatusBadGateway, map[string]string{
			"cause":   "error occur in backend service",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reply)
}
