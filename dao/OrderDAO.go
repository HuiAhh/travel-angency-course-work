package dao

import (
	"travel-agency/core/dto"
	"travel-agency/entity"
	"travel-agency/entity/association"
)

type OrderDAO interface {
	SaveOrUpdate(obj *association.OrderDetailAssociation) (int, error)
	FindList(query *entity.Order, result *dto.PageResult)

	// 模糊查询by order id
	// FindOrderAndDetailList(w, orderId, travelPathId, travelDetailId)
	FindOrderAndDetailList(result *dto.PageResult, orderId *string, travelPathId *string, travelDetailId *string)

	FindOrderAndDetailById(id int) *association.OrderDetailAssociation
}
