package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//首页路由
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "base.html", nil)
}
