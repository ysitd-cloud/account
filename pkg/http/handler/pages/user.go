package pages

import (
	"net/http"

	"code.ysitd.cloud/auth/account/pkg/http/helper"
	"github.com/gin-gonic/gin"
)

func profile(c *gin.Context) {
	helper.RenderAppView(c, http.StatusOK, "", nil)
}
