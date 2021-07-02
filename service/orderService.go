package service

import (
	"travel-agency/core/dto"
	"travel-agency/entity"
	"travel-agency/entity/association"
)

type OrderService interface {
	//插入或者 更新订单
	SaveOrUpdateOrder(order *association.OrderDetailAssociation) (int, error)
	// 变更订单状态
	UpdatePayStatus(order *entity.Order) (int, error)

	//逻辑删除
	DeleteOrder(order *entity.Order) error

	//动态 sql 查询
	FindByCondition(path *entity.Order, page int, size int) dto.PageResult

	//一对一查询order和detail 全部
	// 模糊查询by order id
	//FindWithDetails(page, size, &orderId, &travelPathId, &travelDetailId)
	FindWithDetails(page int, size int, orderId *string, travelPathId *string, travelDetailId *string) dto.PageResult
	GetById(id int) entity.Order

	GetOrderAndDetailById(id int) *association.OrderDetailAssociation
}
