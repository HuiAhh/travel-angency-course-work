package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"travel-agency/util"

	sutil "github.com/lyr-2000/goUtil/util"
)

// file/upload
func UploadFile(c *gin.Context) {
	if f, err := c.FormFile("file"); err == nil {
		newName := sutil.GetRandomString(15)
		newPath := "./static/upload/" + newName + f.Filename
		er := c.SaveUploadedFile(f, newPath)
		log.Println(er)
		util.RenderJson(c, gin.H{
			"location": newPath,
		})

	}else {
		util.RenderJson(c,gin.H{
			"location":"https://ss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=1788874810,942146822&fm=26&gp=0.jpg",
		})
	}
}
