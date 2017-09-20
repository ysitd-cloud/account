package connect

import (
	"crypto/rand"
	"encoding/base64"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/account/http/middlewares"
	"github.com/ysitd-cloud/account/provider"
)

func redirectToOAuth(c *gin.Context) {
	providerID := c.Param("provider")

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	session := middlewares.GetSession(c)
	session.Set(getStateSessionKey(providerID), state)
	session.Save()

	authProvider := provider.GetProvider(providerID)
	config := authProvider.GetConfig()
	url := config.AuthCodeURL(state)

	c.Redirect(http.StatusFound, url)
	c.Abort()
}

func getStateSessionKey(id string) string {
	return "state:" + id
}
