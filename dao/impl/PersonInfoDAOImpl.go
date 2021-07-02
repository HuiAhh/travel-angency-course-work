package impl

import (
	"travel-agency/core/dto"
	"travel-agency/entity"
	"travel-agency/entity/query"
)

type PersonInfoDAOImpl struct {

}
/*
//根据订单信息 查询用户
	findByOrderId(orderId int) []entity.PersonInfo

	//危险操作，逻辑删除， 删除订单下所有的用户信息
	DeleteAllPersonInfoByOrderId(orderId int) (int ,error)

	//动态 sql查询 personInfo信息
	findByExpression(criteria *entity.PersonInfo) (dto.PageResult,[]entity.PersonInfo)


*/



func (*PersonInfoDAOImpl) FindByOrderId(orderId int) (res[]entity.PersonInfo) {
	db.Model(&entity.PersonInfo{}).Where("order_id = ? ",orderId).Find(&res)
	return res
}

//删除订单下 用户的信息
func (*PersonInfoDAOImpl) DeleteAllPersonInfoByOrderId(orderId int) (int ,error) {
	db.Model(&entity.PersonInfo{
		DeleteStatus: true,
		OrderId: orderId,
	})
	return 1,nil
}
//动态 sql查询 personInfo信息
func (*PersonInfoDAOImpl) FindByExpression(criteria *entity.PersonInfo,page int ,size int) (a dto.PageResult,b []entity.PersonInfo) {
	if criteria==nil {
		criteria = &entity.PersonInfo{}
	}
	db.Model(&entity.PersonInfo{}).Where(criteria).Offset((page-1) * size).Limit(size).Find(&b)
	db.Model(&entity.PersonInfo{}).Where(criteria).Count(a.TotalCount)
	a.Result = b
	a.CurrentPage = page
	a.RequestSize = size
	a.CalcTotalPage()
	return a,b
}



func (*PersonInfoDAOImpl ) FindByOrderInfo(personInfo query.PersonInfoSearch,page int ,size int)(res dto.PageResult,list []entity.PersonInfoVO ){

	db.Model(&personInfo).Where(&personInfo).Count(&res.TotalCount);
	db.Where(&personInfo).Limit(size).Offset(size * (page-1)).Find(&list)
	res.RequestSize = size
	res.CurrentPage = page
	res.Result =&list
	res.CalcTotalPage()
	return res, list
}