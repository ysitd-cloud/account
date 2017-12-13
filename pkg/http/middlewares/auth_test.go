package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func createRequestWithAuthHeader(content string) *http.Request {
	header := http.Header{}
	header.Set("Authorization", content)
	return &http.Request{
		Header: header,
	}
}

func TestContainsAuthHeaderNormal(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	req := createRequestWithAuthHeader("Bearer foo")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ContainsAuthHeader(c)

	if c.IsAborted() {
		t.Error("Context should not abort", c)
	}
}

func TestContainsAuthHeaderWrongType(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	req := createRequestWithAuthHeader("bar foo")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ContainsAuthHeader(c)

	if !c.IsAborted() {
		t.Error("Context should abort", c)
	}

	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Error("Wrong status code", resp.StatusCode)
	}
}

func TestContainsAuthHeaderWrongFormat(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	req := createRequestWithAuthHeader("foo")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ContainsAuthHeader(c)

	if !c.IsAborted() {
		t.Error("Context should abort", c)
	}

	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Error("Wrong status code", resp.StatusCode)
	}
}

func TestContainsAuthHeaderWithoutHeader(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	req := &http.Request{}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ContainsAuthHeader(c)

	if !c.IsAborted() {
		t.Error("Context should abort", c)
	}

	resp := w.Result()

	if resp.StatusCode != http.StatusUnauthorized {
		t.Error("Wrong status code", resp.StatusCode)
	}
}
