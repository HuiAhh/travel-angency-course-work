package controller

import (
	"travel-agency/util"

	"github.com/gin-gonic/gin"
)

//daily_income
//每日收支[按周结算]
func DailyIncome(c *gin.Context) {
	//data:{date:[],income:[]} 空处补0
	data:=payService.GetWeeklyIncome()
	util.RenderHtml(c, "daily-income.html", gin.H{
		"data":data.Result,
	})
}

//daily_group
//每日开团
func DailyGroup(c *gin.Context) {
	data:=payService.GetWeeklyDepart()

	util.RenderHtml(c, "daily-group.html", gin.H{
		"data":data.Result,
	})
}
