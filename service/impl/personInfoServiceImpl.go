package impl

import (
	log "github.com/sirupsen/logrus"
	"travel-agency/core/dto"
	"travel-agency/entity"
	"travel-agency/entity/query"
	"travel-agency/entity/resultMap"
	"travel-agency/factory/dao"
	"travel-agency/util"
)

type PersonInfoServiceImpl struct{}

var personInfoDAO = dao.GetPersonInfoDAO()

func (*PersonInfoServiceImpl) FindByOrderId(orderId int) (res []entity.PersonInfo) {
	return personInfoDAO.FindByOrderId(orderId)
}

//删除订单下 用户的信息
func (*PersonInfoServiceImpl) DeleteAllPersonInfoByOrderId(orderId int) (int, error) {
	return personInfoDAO.DeleteAllPersonInfoByOrderId(orderId)
}

//动态 sql查询 personInfo信息
func (*PersonInfoServiceImpl) FindByExpression(criteria *entity.PersonInfo, page int, size int) (a dto.PageResult, b []entity.PersonInfo) {
	return personInfoDAO.FindByExpression(criteria, page, size)
}

//<<<<<<< HEAD
//func (PersonInfoServiceImpl) SaveOne(personInfo *entity.PersonInfo) {
//
//}
//
////删除所有 orderDetailId 的 personInfo
//func (PersonInfoServiceImpl) DeleteAll(orderDetailId int64) {
//
//}
//
////匹配存储
//func (PersonInfoServiceImpl) SaveList(personInfo []entity.PersonInfo) {
//
//}
//
//func (PersonInfoServiceImpl) GetPersonsByOrderDetailId(orderDetailId int) []entity.PersonInfo {
//=======
func (PersonInfoServiceImpl) GetPersonsByOrderId(orderId int) []entity.PersonInfo {
//>>>>>>> 42f77157a29fcdcacd90fbe018a863223cfa4f27
	var persons []entity.PersonInfo
	if orderId<=0{
		return persons
	}
//<<<<<<< HEAD
//	db.Where("delete_status=false and order_id =?",orderDetailId).Find(&persons)
//=======
	db.Where("delete_status=false and order_id =?",orderId).Find(&persons)
//>>>>>>> 42f77157a29fcdcacd90fbe018a863223cfa4f27
	log.Printf("find persons: %v",util.ObjectToStr(persons))
	return persons
}


func ( c *PersonInfoServiceImpl) SavePersonInfo(dto *resultMap.PersonInfoWithOrder) error {
	if dto.OrderId>0 {
		log.Println("删除原信息，",&dto)
		_,err := c.DeleteAllPersonInfoByOrderId(dto.OrderId)
		if err == nil{
			//db.Save(dto.PersonInfoList)
			for _,v:= range dto.PersonInfoList {
				v.OrderId=dto.OrderId
				db.Model(&entity.PersonInfo{}).Save(&v)
			}
		}else{
			log.Error("更新 用户数据异常",err)
			return err
		}
	}else {
		log.Warn("更新失败，没有ID")
	}
	return nil
}

func (c *PersonInfoServiceImpl) FindByOrderInfo(personInfo query.PersonInfoSearch,page int ,size int)(res dto.PageResult,list []entity.PersonInfoVO ) {
	return personInfoDAO.FindByOrderInfo(personInfo,page,size)
}