package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"code.ysitd.cloud/auth/account/third_party/forked/github.com/RangelReale/osin"
	"github.com/gin-gonic/gin"
)

func TestCheckGetUserAccessAccept(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{
		Key:   "user",
		Value: "foo",
	})

	access := &osin.AccessData{
		UserData: "foo",
	}

	c.Set("oauth.access", access)

	CheckGetUserAccess(c)

	if c.IsAborted() {
		t.Error("Context should not abort", c)
	}
}

func TestCheckGetUserAccessForbidden(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = append(c.Params, gin.Param{
		Key:   "user",
		Value: "foo",
	})

	access := &osin.AccessData{
		UserData: "bar",
	}

	c.Set("oauth.access", access)

	CheckGetUserAccess(c)

	if !c.IsAborted() {
		t.Error("Context should abort", c)
	}

	resp := w.Result()
	if resp.StatusCode != http.StatusForbidden {
		t.Error("Wrong status code", resp.StatusCode)
	}
}
