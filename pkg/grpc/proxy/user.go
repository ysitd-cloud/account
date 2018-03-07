package proxy

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"code.ysitd.cloud/grpc/schema/account/actions"
	"github.com/tonyhhyip/vodka"
)

func (p *proxy) getUserInfo(c *vodka.Context) {
	service := p.service

	username, err := vodka.String(c.UserValue("username"))
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"cause":   "Fail to parse username from url",
			"message": err.Error(),
		})
	}

	reply, err := service.GetUserInfo(c, &actions.GetUserInfoRequest{
		Username: username,
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

func (p *proxy) validateUserPassword(c *vodka.Context) {
	service := p.service

	var req actions.ValidateUserRequest
	content, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {

	}
	json.Unmarshal(content, &req)

	reply, err := service.ValidateUserPassword(c, &req)

	if err != nil {
		c.JSON(http.StatusBadGateway, map[string]string{
			"cause":   "error occur in backend service",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, reply)
}
