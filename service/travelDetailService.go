package service

import (
	"travel-agency/core/dto"
	"travel-agency/entity"
	"travel-agency/entity/resultMap"
)

type TravelDetailService interface {
	// //插入或者 更新订单
	// SaveOrUpdateOrder(order *entity.Order ) (int,error)
	// //逻辑删除
	// DeleteOrder(order *entity.Order) (int,error)

	// //查询
	// SelectOrder(query *entity.Order) []entity.Order

	//更新或者保存旅游团信息

	SaveOrUpdateTravelDetail(detail *entity.TravelPathDetail) (int, error)

	//分页查询搜索 出 对应 的 旅游团信息
	QueryDetailInfoByPathId(detail *entity.TravelPathDetail, page int, size int) (dto.PageResult, []resultMap.TravelDetailWithOrderPersonCount)
	//SaveOne(personInfo *entity.PersonInfo)
	////删除所有 orderDetailId 的 personInfo
	//DeleteAll(orderDetailId int64)
	////匹配存储
	//SaveList(personInfo []entity.PersonInfo)

	// @Author: HuiAhh
	// @Description: get by travel detail id
	GetById(id int) *entity.TravelPathDetail
}
