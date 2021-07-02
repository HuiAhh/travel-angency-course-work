package util

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func StrToJson(c *gin.Context, obj interface{}) error {
	err := c.ShouldBindJSON(obj)
	log.Printf("json = %v", ToStringJSON((obj)))
	return err
}

func ParseJson(txt string, obj interface{}) error {
	er := json.Unmarshal([]byte(txt), &obj)
	return er
}

func RenderJson(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, obj)
}



//对象转字符串
func ObjectToStr(obj interface{}) string {
	if obj == nil {
		return ""
	}
	res,_ := json.Marshal(obj)
	return string(res)
}

func ToStringJSON(obj interface{}) string {
	if obj == nil {
		return ""
	}
	parsed, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(parsed)
}
