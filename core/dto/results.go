package dto

import (
	"github.com/gin-gonic/gin"
	"travel-agency/core/enums"
	"travel-agency/util"
)

////通用返回结果
//type Result struct {
//	//result结果
//	result interface{}
//	code enums.ResultStatus
//	message string
//}

type PageResult struct {
	TotalCount  int //数据总数
	TotalPage   int //总页数
	CurrentPage int //当前页

	RequestSize int         //请求多少条
	Result      interface{} //携带的返回 值
	ExtraMsg    string      //其他信息

}

func (s *PageResult) String() string {
	return util.ObjectToStr(s)
}

//func (s PageResult) Calc() {
//
//	//总页数
//	s.TotalPage = ((s.TotalCount - 1) / s.RequestSize) + 1
//	fmt.Println(s.TotalPage)
//}
func (s *PageResult) CalcTotalPage() int {
	//总页数
	s.TotalPage = ((s.TotalCount - 1) / s.RequestSize) + 1
	return s.TotalPage
}
func Ok(res interface{}) interface{} {
	return gin.H{
		"result": res,
		"code":   enums.OK,
	}

}

func Fail(res interface{}, errMsg string) interface{} {
	return gin.H{

		"result": res,
		"code":   enums.Error,
		"errMsg": errMsg,
	}
}

func FailWithErr(res interface{}, customMsg string, err error) interface{} {
	var h gin.H = Fail(res, customMsg).(gin.H)
	h["err"] = err
	return h
}
