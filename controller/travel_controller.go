package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"html/template"
	"net/http"
	"travel-agency/core/dto"
	"travel-agency/entity"
	"travel-agency/factory/service"

	valid "travel-agency/factory/validator"
	//"travel-agency/factory/dao"
	"travel-agency/util"
)

//获取校验器
var validator = valid.GetValidator()
var travelService = service.GetTravelPathService()

//travel_table
//表格预览     旅游路径
func TravelTable(c *gin.Context) {
	page, size := util.GetPageAndSize(c)
	keyword := c.Query("keyword")
	res := travelService.FindByCondition(&entity.TravelPath{
		TravelPathName: &keyword,
	}, page, size)
	//c.HTML(http.StatusOK, "travel_table.html", res)
	log.Println(&res)
	util.RenderHtml(c, "travel_table.html", gin.H{
		"res": res,
		"keyword":keyword, //搜索条件
	})
}

//travel_edit
//编辑和添加旅游路径  【修改 之类 的】
func FormEdit(c *gin.Context) {
	//TODO: 如果 id 不为null ，就跳到编辑页面，否则跳到新增页面
	//travelPathId := c.Query("id")
	//if  util.IsNotBlank(travelPathId) {
	//	//添加页面
	//	//c.HTML(http.StatusOK, "travel_table.html", nil)
	//}

	if Id, err := util.ToInt(c.Query("id")); err != nil || Id <= 0 {
		c.HTML(http.StatusOK, "travel_edit.html", gin.H{
			"res": nil,
		})
	} else {
		out := travelService.GetById(Id)
		//log.Printf("res = %s", *out.TravelPathContent)
		//根据 ID 更新
		util.RenderHtml(c, "travel_change.html", gin.H{
			"entity":  out,
			"content": template.HTML(*out.TravelPathContent),
		})
	}

	// 查询出来
	//id,err := util.ToInt(travelPathId)

}

//通过 JSON 更新
//update_travel_path
func UpdateTravelPath(c *gin.Context) {
	json := entity.TravelPath{}

	_ = util.StrToJson(c, &json)
	if  json.TravelPathId > 0 {
		//更新上下架状态
		_, _ = travelService.UpdateActiveStatus(&json)
	}
	util.RenderJson(c, dto.Ok(nil))
}
//del_travel_path
func DeleteTravelPath(c *gin.Context) {
	if id,err := util.ToInt(c.Query("id"));err !=nil {
		travelService.DeleteById(&entity.TravelPath{TravelPathId: id})
		util.RenderJson(c,dto.Ok(nil))
	}else{
		util.RenderJson(c,dto.Fail(c,"删除失败"))
	}




}

//travel_path/save
//保存
func SaveTravelPathJson(c *gin.Context) {
	//json := entity.TravelPath{}
	////TODO: 保存 ajax传输的 json
	//json.Unmarshal()
	json := entity.TravelPath{}

	_ = util.StrToJson(c, &json)

	log.Info("json ::", &json)
	if err := validator.Struct(json); err != nil || util.IsBlank(json.TravelDescription) {
		log.Info("参数校验异常  %v, %v", err)
		util.RenderJson(c, dto.FailWithErr(&json, "参数校验失败,输入不符合格式", err))
	} else {
		if json.TravelPathContent != nil {
			newContent := util.ReplaceHtml(*json.TravelPathContent)
			//去掉脚本标签
			json.TravelPathContent = &newContent
		}

		if r, err := travelService.SaveOrUpdate(&json); err != nil {
			util.RenderJson(c, dto.FailWithErr(json, "保存出现错误", err))
		} else if r <= 0 {
			util.RenderJson(c, dto.Fail(&json, "保存失败"))
		} else {
			util.RenderJson(c, dto.Ok(&json))
		}
	}

}

//travel_path/delete
func DeleteByTravelPathId(c *gin.Context) {
	if Id,err := util.ToInt(c.Query("id")); err ==nil  && Id >0 {
		err = travelService.DeleteById(&entity.TravelPath{TravelPathId: Id})
		if err ==nil {
			util.RenderJson(c, dto.Ok(nil))
		}else {
			util.RenderJson(c,dto.Fail(nil,"删除失败，操作异常"))
		}

	}else {
		util.RenderJson(c,dto.Fail(nil,"删除失败，没有这条数据"))

	}



}


// 预览旅游路线
//  travelPath/preview
func PreviewTravelPathId( c *gin.Context) {
	if Id,err := util.ToInt(c.Query("id")); err ==nil  && Id >0 {
		e := travelService.GetById(Id)
		log.Println(&e)
		if e.TravelPathContent ==nil {
			s := ""
			e.TravelPathContent = &s
		}
		//查询旅游路线，并且 展示到页面上
		util.RenderHtml(c,"preview_travelPath.html",gin.H{
			"res":e,
			"content":template.HTML(*e.TravelPathContent ),
		})
		//err = travelService.DeleteById(&entity.TravelPath{T1ravelPathId: Id})
		//if err ==nil {
		//	util.RenderJson(c, dto.Ok(nil))
		//}else {
		//	util.RenderJson(c,dto.Fail(nil,"删除失败，操作异常"))
		//}

	}else {
		//util.RenderJson(c,dto.Fail(nil,"删除失败，没有这条数据"))
		util.RenderHtml(c,"404.html",nil)
	}
}