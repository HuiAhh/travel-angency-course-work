package dao

import (
	"travel-agency/core/dto"
)

type PayDAO interface {
	// 模糊查询order id
	FindNonPaidOrders(result *dto.PageResult, keyword *string)
	FindPaidOrders(result *dto.PageResult, keyword *string)

	//根据id查
	FindNonAndPaidOrderById(result *dto.PageResult, id int)

}
