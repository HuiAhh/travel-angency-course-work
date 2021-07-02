package util

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetPageAndSize(c *gin.Context) (int, int) {
	page := c.Query("page")
	size := c.Query("size")
	pageNum, err := ToInt(page)
	if err != nil {
		log.Warn("page  ", err)
		pageNum = 1
	}
	SizeNum, err := ToInt(size)
	if err != nil {
		SizeNum = 5
		log.Warn("size ", size)
	}
	return pageNum, SizeNum
}
