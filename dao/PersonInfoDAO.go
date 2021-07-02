package dao

import (
	"travel-agency/core/dto"
	"travel-agency/entity"
	"travel-agency/entity/query"
)

type PersonInfoDAO interface {
	//根据订单信息 查询用户
	FindByOrderId(orderId int) []entity.PersonInfo

	//危险操作，逻辑删除， 删除订单下所有的用户信息
	DeleteAllPersonInfoByOrderId(orderId int) (int ,error)

	//动态 sql查询 personInfo信息
	FindByExpression(criteria *entity.PersonInfo,page int,size int) (dto.PageResult,[]entity.PersonInfo)

	FindByOrderInfo(personInfo query.PersonInfoSearch,page int ,size int)(res dto.PageResult,list []entity.PersonInfoVO )

}
