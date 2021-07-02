package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"html/template"
	"travel-agency/core/dto"
	"travel-agency/core/exception"
	"travel-agency/entity"
	f "travel-agency/factory/service"
	"travel-agency/util"
)

// 旅游团信息 CRUD 相关
var detailService = f.GetTravelPathDetailService()

//保存和更新旅游团

// @Router /travel_detail/add [post]
func AddTravelPathDetail(c *gin.Context) {
	//util.RenderHtml(c, "travel_add_detail.html" ,nil)
	w := entity.TravelPathDetail{}
	if err := util.StrToJson(c, &w); err == nil {
		//var timeErr error
		if w.BeginTime !=nil && w.EndTime!=nil && w.BeginTime.AsInt() >= w.EndTime.AsInt() {
			log.Warn("结束时间大于开始时间")
			errMsg := "结束时间大于开始时间"
			timeErr := exception.TimeIsError(&errMsg)
			util.RenderJson(c, dto.Fail(timeErr, "校验失败，结束时间大于开始时间"))
			return
		}

		validErr := validator.Struct(w)
		if   validErr == nil {
			// 处理
			res, er := detailService.SaveOrUpdateTravelDetail(&w)
			if er == nil && res > 0 {
				util.RenderJson(c, dto.Ok(nil))
			} else {
				util.RenderJson(c, dto.Fail(er, "更新失败-数据库异常"))
			}
		} else {
			util.RenderJson(c, dto.Fail(validErr, "校验失败，字段有问题"))
		}

	} else {
		util.RenderJson(c, dto.Fail(err, "json 解析失败了"))
	}

}

// @Router /travel_detail/query [get]
//@Description 查询接口

func QueryTravelDetail(c *gin.Context) {
	w := entity.TravelPathDetail{}
	//分页查询
	page, size := util.GetPageAndSize(c)
	keyword := c.Query("teamName")
	id, _ := util.ToInt(c.Query("id"))
	w.TeamName = &keyword
	w.TravelPathId = &id
	pageResult, list := detailService.QueryDetailInfoByPathId(&w, page, size)
	log.Println(&pageResult)
	util.RenderHtml(c, "travel_detail.html", gin.H{
		"list":       list,
		"pageResult": pageResult,
		"keyword":    &keyword,
		"id":         id,
	})

}

//travel_detail/detail?id=
func PreviewTravelDetail(c *gin.Context) {
	if id,err := util.ToInt(c.Query("id")); err ==nil  {
		//查询
		w := detailService.GetById(id)
		ptr := w.TravelPathId
		pathId :=0
		if ptr !=nil {
			pathId  = *ptr
		}
		path := travelService.GetById(pathId)
		html := path.TravelPathContent
		if html==nil {
			html = new(string)
		}
		util.RenderHtml(c,"preview_travelPath_detail.html",gin.H{
			"res":w,
			"path":path,
			"html":template.HTML(*html),
		})
	}
}
//travel_detail/edit
func EditTravelDetail(c * gin.Context) {

	if id,err:=util.ToInt(c.Query("id")) ; err==nil && id >0{
		w := entity.TravelPathDetail{}
		w.TravelDetailId = id
		result := detailService.GetById(id)
		util.RenderHtml(c,"travel_detail_edit.html",gin.H{
			"result":result,
		})

	}
}

//func UpdateTravelDetail(c *gin.Context) {
//	w := entity.TravelPathDetail{}
//	if err:=util.StrToJson(c, &w) ; err==nil {
//		validErr := validator.Struct(w)
//		if validErr ==nil {
//			// 处理
//			res,er := detailService.SaveOrUpdateTravelDetail(&w)
//			if er == nil && res >0 {
//				util.RenderJson(c,dto.Ok(nil))
//			}else{
//				util.RenderJson(c,dto.Fail(er,"更新失败-数据库异常"))
//			}
//		}else {
//			util.RenderJson(c,dto.Fail(validErr,"校验失败，字段有问题"))
//		}
//
//	}
//
//}
