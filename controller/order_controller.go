package controller

import (
	"html/template"
	"net/http"
	"travel-agency/core/dto"
	"travel-agency/entity"
	"travel-agency/entity/association"

	"travel-agency/factory/service"
	"travel-agency/util"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var orderService = service.GetOrderService()

// order_edit
// 编辑订单  【修改 之类 的】
// copy func-FormEdit
func OrderEditPage(c *gin.Context) {

	if Id, err := util.ToInt(c.Query("id")); err != nil || Id <= 0 {
		c.HTML(http.StatusOK, "order_edit.html", gin.H{
			"res": nil,
		})

	} else {
		out := orderService.GetOrderAndDetailById(Id)
		var content string = ""
		if out.OrderDetailDescription != nil {
			content = *out.OrderDetailDescription
		}
		util.RenderHtml(c, "order_change.html", gin.H{
			"entity": out,

			"content": template.HTML(content),
		})
	}

	// 查询出来
	//id,err := util.ToInt(travelPathId)

}

//order_table
//表格预览     订单列表
func OrderTable(c *gin.Context) {
	page, size := util.GetPageAndSize(c)
	// order id 作精确搜索条件
	orderId := c.Query("orderId")
	travelPathId := c.Query("travelPathId")
	travelDetailId := c.Query("travelDetailId")

	res := orderService.FindWithDetails(page, size, &orderId, &travelPathId, &travelDetailId)
	log.Printf("orders and details: %v", util.ObjectToStr(res))
	util.RenderHtml(c, "order_table.html", gin.H{
		"res":            res,
		"orderId":        orderId,
		"travelPathId":   travelPathId,
		"travelDetailId": travelDetailId,
	})
}

// JSON
//update_order
// 更新订单 支付状态
func UpdateOrderPayStatus(c *gin.Context) {
	json := entity.Order{}

	_ = util.StrToJson(c, &json)
	if json.OrderId > 0 {
		//更新支付状态
		_, _ = orderService.UpdatePayStatus(&json)
	}
	util.RenderJson(c, dto.Ok(nil))
}

//order/save
//保存
func SaveOrderJson(c *gin.Context) {
	//json := entity.TravelPath{}
	////TODO: 保存 ajax传输的 json
	//json.Unmarshal()
	json := association.OrderDetailAssociation{}

	_ = util.StrToJson(c, &json)

	log.Info("json ::", util.ObjectToStr(json))
	if err := validator.Struct(json); err != nil || util.IsBlank(json.OrderDetailDescription) {
		log.Info("参数校验异常  %v, %v", err)
		util.RenderJson(c, dto.FailWithErr(&json, "参数校验失败,输入不符合格式", err))
	} else {

		if *json.ChildCount > *json.PersonCount {
			util.RenderJson(c, dto.FailWithErr(&json, "儿童人数不能大于总人数", err))
			return
		}

		if json.OrderDetailDescription != nil {
			newContent := util.ReplaceHtml(*json.OrderDetailDescription)
			//去掉脚本标签
			json.OrderDetailDescription = &newContent
		}

		if r, err := orderService.SaveOrUpdateOrder(&json); err != nil {
			util.RenderJson(c, dto.FailWithErr(json, "保存出现错误", err))
		} else if r <= 0 {
			util.RenderJson(c, dto.Fail(&json, "保存失败"))
		} else {
			util.RenderJson(c, dto.Ok(&json))
		}
	}

}

//order/delete
//根据ID逻辑删除
func DeleteByOrderId(c *gin.Context) {
	if Id, err := util.ToInt(c.Query("id")); err == nil && Id > 0 {
		err = orderService.DeleteOrder(&entity.Order{OrderId: Id})
		if err == nil {
			util.RenderJson(c, dto.Ok(nil))
		} else {
			util.RenderJson(c, dto.Fail(nil, "删除失败，操作异常"))
		}

	} else {
		util.RenderJson(c, dto.Fail(nil, "删除失败，没有这条数据"))

	}
}
