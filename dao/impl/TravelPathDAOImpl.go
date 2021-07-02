package impl

import (
	"errors"
	"travel-agency/core/dto"
	"travel-agency/entity"
	"travel-agency/util"

	log "github.com/sirupsen/logrus"

	dbFactory "travel-agency/factory/db"

	"github.com/jinzhu/gorm"
)

type TravelPathDAOImpl struct {
}

//func (t TravelPathDAOImpl) SaveOrUpdate(obj *entity.TravelPath) (int, error) {
//	panic("implement me")
//}

var db *gorm.DB = dbFactory.GetGormDB()

func (*TravelPathDAOImpl) SaveOrUpdate(obj *entity.TravelPath) (int, error) {
	if obj == nil {

		return 0, errors.New("空指针异常")

	}
	log.Printf("args ->  xx %v", util.ObjectToStr(obj))
	//db.Save(obj)
	if  obj.TravelPathId <= 0 {
		//save
		db.Create(obj)

	} else {
		//update
		db.Model(entity.TravelPath{}).Where("travel_path_id = ?", obj.TravelPathId).Updates(obj)
	}

	return 1, nil

}

func (*TravelPathDAOImpl) FindList(query *entity.TravelPath, result *dto.PageResult) {
	if query == nil || result == nil {
		return
	}
	var res []entity.TravelPath
	//var res []entity.TravelPath
	keyword := query.TravelPathName
	query.TravelPathName = nil
	db.
		Where("delete_status = ?  and travel_path_name like concat('%',?,'%')",false, keyword).

		Offset((result.CurrentPage - 1) * result.RequestSize).
		Limit(result.RequestSize).
		Order("active_status desc").
		Find(&res)
	db.Model(query).
		Where("delete_status = ?  " +
		"and " +
		"travel_path_name like concat('%',?,'%')",false, keyword).Count(&result.TotalCount)
	//总页数
	result.TotalPage = ((result.TotalCount - 1) / result.RequestSize) + 1

	result.Result = res
	//db.Model(query).Find(query).Limit()
	return

}
