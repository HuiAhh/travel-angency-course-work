package service

import (
	"travel-agency/core/dto"
	"travel-agency/entity"
)

type TravelService interface {
	//保存 或者插入， 如果 id 为空 就插入，否则保存
	SaveOrUpdate(path *entity.TravelPath) (int, error)
	//根据ID 删除
	DeleteById(path *entity.TravelPath) error
	//动态 sql 查询
	FindByCondition(path *entity.TravelPath, page int, size int) dto.PageResult
	//更新上下架状态
	UpdateActiveStatus(path *entity.TravelPath) (int, error)

	GetById(id int) entity.TravelPath
}

//func SaveOrUpdate( path *entity.TravelPath) (int,error) {
//	return 0,nil
//}
