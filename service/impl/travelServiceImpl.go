package impl

import (
	"travel-agency/core/dto"
	"travel-agency/core/exception"
	"travel-agency/entity"
	"travel-agency/factory/dao"
	dbFactory "travel-agency/factory/db"
	"travel-agency/util"

	log "github.com/sirupsen/logrus"
)

type TravelServiceImpl struct {
}
//
var travelPathDAO = dao.GetTravelPathDao()
var db = dbFactory.GetGormDB()

func (cur *TravelServiceImpl) DeleteById(path *entity.TravelPath) error {
	//panic("implement me")
	if path != nil && path.TravelPathId > 0 {
		db.Model(path).Update("delete_status", true)
	}
	return nil
}

func (cur *TravelServiceImpl) FindByCondition(path *entity.TravelPath, page int, size int) dto.PageResult {
	//panic("implement me")
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 5
	}
	w := &dto.PageResult{
		RequestSize: size,
		CurrentPage: page,
	}
	s := false
	path.DeleteStatus = &s
	travelPathDAO.FindList(path, w)

	return *w
	//travelPathDAO
}

/*
//保存 或者插入， 如果 id 为空 就插入，否则保存
	SaveOrUpdate(path *entity.TravelPath)(int,error)
	//根据ID 删除
	DeleteById(path *entity.TravelPath) error
	//动态 sql 查询
	FindByCondition(path *entity.TravelPath,page int ,size int) []entity.TravelPath

*/

//更新 或者 插入一条数据

func (cur *TravelServiceImpl) SaveOrUpdate(path *entity.TravelPath) (int, error) {
	log.Printf("args -> = %v", path)
	if path == nil {
		return 0, exception.Npe(nil)

	}
	//save
	return travelPathDAO.SaveOrUpdate(path)

}

func (cur *TravelServiceImpl) UpdateActiveStatus(path *entity.TravelPath) (int, error) {
	log.Printf("request = %v", util.ToStringJSON(*path))
	// 2020-5-24 12:26AM fix update: false->true 要传个activeStatus的指向进去
	if path != nil {
		db.Model(path).
			Update("active_status", *path.ActiveStatus)
	}
	return 1, nil

}

func (cur *TravelServiceImpl) GetById(id int) (res entity.TravelPath) {
	log.Printf("find byid = %d", id)

	db.Model(&res).Where("travel_path_id = ?", id).
		First(&res)
	return res
}
