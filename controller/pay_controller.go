package controller

import (
	"log"
	"net/http"
	"travel-agency/factory/service"
	"travel-agency/util"

	"github.com/gin-gonic/gin"
)

var (
	payService        = service.GetPayService()
	personInfoService = service.GetPersonInfoService()
)

//pay_fee_table
//获取订金费率 手续费率表
func PayFeeTable(c *gin.Context) {
	res := payService.GetAll()
	util.RenderHtml(c, "rate-table.html", gin.H{
		"res": res,
	})
}

//non_paid_table
//未支付旅游团订单 表
func NonPaidTable(c *gin.Context) {
	page, size := util.GetPageAndSize(c)
	// order id 作模糊搜索条件
	keyword := c.Query("keyword")
	res := payService.GetNonPaidOrders(page, size, &keyword)
	log.Printf("non paid orders: %v", util.ObjectToStr(res))
	util.RenderHtml(c, "non-paid-table.html", gin.H{
		"res":     res,
		"keyword": keyword,
	})
}

//paid_table
//已支付旅游团订单 表
func PaidTable(c *gin.Context) {
	page, size := util.GetPageAndSize(c)
	// order id 作模糊搜索条件
	keyword := c.Query("keyword")
	res := payService.GetPaidOrders(page, size, &keyword)
	log.Printf("non paid orders: %v", util.ObjectToStr(res))
	util.RenderHtml(c, "paid-table.html", gin.H{
		"res":     res,
		"keyword": keyword,
	})
}

//pay/doc
//生成交款单
func PayDocument(c *gin.Context) {
	if Id, err := util.ToInt(c.Query("id")); err != nil || Id <= 0 {
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
	} else {
		res := payService.GetNonAndPaidOrderById(Id)

		util.RenderHtml(c, "pay-doc.html", gin.H{
			"res": res,
		})
	}
}

//pay/confirmation
//生成确认书
func PayConfirmation(c *gin.Context) {
	if Id, err := util.ToInt(c.Query("id")); err != nil || Id <= 0 {
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
	} else {
		res := payService.GetNonAndPaidOrderById(Id)

		orderDetail := orderService.GetOrderAndDetailById(Id)

		// 查旅行团游客人员
		persons := personInfoService.GetPersonsByOrderId(Id)

		//旅行团描述
		var teamDescription *string
		if orderDetail.TravelDetailId != nil {
			team := detailService.GetById(*orderDetail.TravelDetailId)
			teamDescription = team.TeamDescription
		}

		//订单描述
		description := orderDetail.OrderDetailDescription
		util.RenderHtml(c, "confirmation-doc.html", gin.H{
			"res":             res,
			"persons":         persons,
			"description":     description,
			"teamDescription": teamDescription,
		})
	}

}
