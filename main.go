package main

import (
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"

	"database/sql"
	"os"
	"github.com/ory/osin-storage/storage/postgres"
	"github.com/RangelReale/osin"
	"gopkg.in/gin-gonic/gin.v1"
)

func handleAuthorize(server *osin.Server) (handlerFunc gin.HandlerFunc) {
	return func (c *gin.Context) {
		resp := server.NewResponse()
		defer resp.Close()

		if ar := server.HandleAuthorizeRequest(resp, c.Request); ar != nil {
			c.Next()
			return
		}
		if resp.IsError && resp.InternalError != nil {
			c.AbortWithError(500, resp.InternalError)
		}
		osin.OutputJSON(resp, c.Writer, c.Request)
		c.Abort()
	}
}

func main() {
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}

	store := postgres.New(db)
	config := osin.NewServerConfig()
	config.AllowedAccessTypes = osin.AllowedAccessType{osin.AUTHORIZATION_CODE, osin.REFRESH_TOKEN}
	config.ErrorStatusCode = 400
	server := osin.NewServer(config, store)

	app := gin.Default()

	app.GET("/authorize", handleAuthorize(server))

	app.Run(":" + os.Getenv("PORT"))
}
