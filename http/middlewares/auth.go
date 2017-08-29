package middlewares

import (
	"net/http"
	"strings"

	"github.com/RangelReale/osin"
	"github.com/ysitd-cloud/judge-go-client"
	"gopkg.in/gin-gonic/gin.v1"
)

func BearerToken(c *gin.Context) {

	if c.MustGet("authorization.type").(string) != "bearer" {
		c.Next()
		return
	}

	server := c.MustGet("osin.server").(*osin.Server)
	token := c.MustGet("authorization.value").(string)

	if access, err := server.Storage.LoadAccess(token); err != nil {
		c.AbortWithStatus(http.StatusForbidden)
	} else {
		c.Set("oauth.access", access)
		c.Next()
	}
}

func JudgeToken(action, resource string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.MustGet("authorization.type").(string) != "judge" {
			c.Next()
			return
		}
		token := c.MustGet("authorization.value").(string)

		subjectHeader := c.GetHeader("X-Client")
		pieces := strings.Split(subjectHeader, ":")
		if len(pieces) != 2 {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		subjectType := judge.SubjectType(pieces[0])
		subjectID := pieces[1]
		client :=
			c.MustGet("judge").(judge.Client)
		subject := judge.NewSubject(subjectID, subjectType)
		result, reason, errors := client.Judge(subject, action, resource, token)
		if len(errors) > 0 {
			c.AbortWithError(http.StatusBadGateway, errors[0])
			return
		}

		if result {
			c.Next()
			return
		}

		if reason == "forbidden" {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

func ContainsJudgeHeader(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	pieces := strings.Split(authHeader, " ")
	if len(pieces) != 2 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	authType := strings.ToLower(pieces[0])
	if authType != "judge" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Set("authorization.type", authType)
	c.Set("authorization.value", pieces[1])

	c.Next()
}

func ContainsAuthHeader(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	pieces := strings.Split(authHeader, " ")
	if len(pieces) != 2 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	authType := strings.ToLower(pieces[0])

	if authType != "bearer" && authType != "judge" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Set("authorization.type", authType)
	c.Set("authorization.value", pieces[1])

	c.Next()
}
