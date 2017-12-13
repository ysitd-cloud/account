package connect

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/account/pkg/http/helper"
)

func listConnect(c *gin.Context) {
	helper.RenderAppView(c, http.StatusOK, "connect", nil)
}
