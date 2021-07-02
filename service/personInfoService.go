package service

import (
	"travel-agency/core/dto"
	"travel-agency/entity"
	"travel-agency/entity/query"
	"travel-agency/entity/resultMap"
)

type PersonInfoService interface {
	// //插入或者 更新订单
	//// SaveOrUpdateOrder(order *entity.Order ) (int,error)
	//// //逻辑删除
	//// DeleteOrder(order *entity.Order) (int,error)
	//
	//// //查询
	//// SelectOrder(query *entity.Order) []entity.Order
	//
	//SaveOne(personInfo *entity.PersonInfo)
	////删除所有 orderDetailId 的 personInfo
	//DeleteAll(orderDetailId int64)
	////匹配存储
	//SaveList(personInfo []entity.PersonInfo)
	//
	//GetPersonsByOrderDetailId(orderDetailId int) []entity.PersonInfo

	//根据订单信息 查询用户
	FindByOrderId(orderId int) []entity.PersonInfo

	//危险操作，逻辑删除， 删除订单下所有的用户信息
	DeleteAllPersonInfoByOrderId(orderId int) (int ,error)

//<<<<<<< HEAD
	//动态 sql查询 personInfo信息
	FindByExpression(criteria *entity.PersonInfo,page int,size int) (dto.PageResult,[]entity.PersonInfo)



	SavePersonInfo(dto *resultMap.PersonInfoWithOrder) error
//=======
	GetPersonsByOrderId(orderId int) []entity.PersonInfo
//>>>>>>> 42f77157a29fcdcacd90fbe018a863223cfa4f27

	FindByOrderInfo(personInfo query.PersonInfoSearch,page int ,size int)(res dto.PageResult,list []entity.PersonInfoVO)
}
