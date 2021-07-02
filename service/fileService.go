package service

import "github.com/gin-gonic/gin"

type FileService interface {
	//crud
	UploadFile(c *gin.Context) (string, error)

	//对detail的crud
}
