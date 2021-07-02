package service

import (
	"travel-agency/core/dto"
	"travel-agency/entity"
)

//CRUD   调汇率
type PayService interface {
	GetAll() []entity.MoneyShouldPayFee

	// 模糊查询order id
	GetNonPaidOrders(page int, size int, keyword *string) dto.PageResult
	GetPaidOrders(page int, size int, keyword *string) dto.PageResult

	//交款单
	GetNonAndPaidOrderById(id int) dto.PageResult

	//周收入报表
	GetWeeklyIncome() dto.PageResult

	//周开团数量
	GetWeeklyDepart() dto.PageResult

}
