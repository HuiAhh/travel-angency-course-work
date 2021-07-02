package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//封装 gin 的 输出方法
func RenderHtml(c *gin.Context, template string, obj interface{}) {
	c.HTML(http.StatusOK, template, obj)
}
