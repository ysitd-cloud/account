package connect

import (
	"context"
	"net/http"

	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/account/http/middlewares"
	"github.com/ysitd-cloud/account/oauth"
)

func oauthCallback(c *gin.Context) {
	providerID := c.Param("provider")

	session := middlewares.GetSession(c)
	stateKey := getStateSessionKey(providerID)
	if !session.Exists(stateKey) {
		c.Redirect(http.StatusBadRequest, "/connect")
		return
	}

	state := session.Get(getStateSessionKey(providerID)).(string)
	if state != c.Query("state") {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	code := c.Query("code")
	if len(code) == 0 {
		c.AbortWithStatus(http.StatusPreconditionFailed)
		return
	}

	authProvider := oauth.GetProvider(providerID)
	config := authProvider.GetConfig()
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		c.AbortWithError(http.StatusBadGateway, err)
		return
	}

	if !token.Valid() {
		c.AbortWithStatus(http.StatusFailedDependency)
		return
	}

	username := session.Get("username").(string)

	id, err := authProvider.GetUserID(token)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	query := `
	INSERT INTO user_connect (username, provider, user_id) VALUES ($1, $2, $3)
	ON CONFLICT (username, provider) DO UPDATE SET user_id = $3
	`
	db := c.MustGet("db").(*sql.DB)
	defer db.Close()

	_, err = db.Exec(query, username, providerID, id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	} else {
		c.Redirect(http.StatusFound, "/connect")
		c.Abort()
	}
}
