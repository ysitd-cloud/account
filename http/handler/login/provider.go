package login

import (
	"crypto/rand"
	"encoding/base64"

	"net/http"

	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/account/http/middlewares"
	"github.com/ysitd-cloud/account/model"
	"github.com/ysitd-cloud/account/provider"
)

func redirectToOAuth(c *gin.Context) {
	providerID := c.Param("provider")

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	session := middlewares.GetSession(c)
	session.Set(getStateSessionKey(providerID), state)
	session.Set("provider:usage", "login")
	session.Set("next:url", c.DefaultQuery("next", "/"))
	session.Save()

	authProvider := oauth.GetProvider(providerID)
	config := authProvider.GetConfig()
	url := config.AuthCodeURL(state)

	c.Redirect(http.StatusFound, url)
	c.Abort()
}

func getStateSessionKey(id string) string {
	return "state:" + id
}

func oauthLoginCallback(c *gin.Context) {
	providerID := c.Param("provider")

	session := middlewares.GetSession(c)
	state, convert := session.Get(getStateSessionKey(providerID)).(string)
	if !convert {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if state != c.Query("state") {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	code := c.Query("code")
	if len(code) == 0 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	authProvider := oauth.GetProvider(providerID)
	config := authProvider.GetConfig()
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if !token.Valid() {
		c.AbortWithStatus(http.StatusFailedDependency)
		return
	}

	id, err := authProvider.GetUserID(token)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	query := `
	SELECT username, display_name, email, avatar_url FROM users
	INNER JOIN user_connect
	WHERE provider = $1 AND user_id = $2
	`
	db := c.MustGet("db").(*sql.DB)

	row := db.QueryRow(query, providerID, id)
	user, err := model.LoadFromRow(row)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	session.Set("username", user.Username)
	session.Set("email", user.Email)
	session.Set("avatar_url", user.AvatarUrl)
	session.Set("display_name", user.DisplayName)
	session.Save()

	c.Redirect(http.StatusFound, session.Get("next:url").(string))
}
