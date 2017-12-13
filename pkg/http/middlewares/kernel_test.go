package middlewares

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/tonyhhyip/go-di-container"
	providers "github.com/ysitd-cloud/account/pkg/kernel"
)

func TestBindKernel(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	BindKernel(c)

	kernel := c.MustGet("kernel").(container.Kernel)

	if kernel != providers.Kernel {
		t.Error("Bind incorrect kernel")
	}
}
