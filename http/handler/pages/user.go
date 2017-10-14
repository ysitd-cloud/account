package pages

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ysitd-cloud/account/http/helper"
)

func profile(c *gin.Context) {
	helper.RenderAppView(c, http.StatusOK, "", nil)
}

func modifiedPassword(c *gin.Context) {
	helper.RenderAppView(c, http.StatusOK, "password", nil)
}
