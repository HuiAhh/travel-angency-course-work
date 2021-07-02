package impl

import (
	"travel-agency/core/dto"
	"travel-agency/entity"
	"travel-agency/entity/resultMap"
	daofactory "travel-agency/factory/dao"
)

type TravelPathDetailServiceImpl struct {
}

var travelPathDetailDAO = daofactory.GetTravelPathDetailDao()

//var db = dbFactory.GetGormDB()

func (*TravelPathDetailServiceImpl) SaveOrUpdateTravelDetail(detail *entity.TravelPathDetail) (int, error) {
	return travelPathDetailDAO.SaveTravelPathDetail(detail)

}

//根据 路径ID 查询旅游团
func (*TravelPathDetailServiceImpl) QueryDetailInfoByPathId(detail *entity.TravelPathDetail, page int, size int) (dto.PageResult, []resultMap.TravelDetailWithOrderPersonCount) {
	if detail.TravelPathId != nil && *detail.TravelPathId > 0 {
		// 根据旅游团信息查询
		//db.Raw("select * from t_")
		return travelPathDetailDAO.FindListTravelPathDetail(detail, page, size)
	}
	return dto.PageResult{}, nil
}

// @Author: HuiAhh
// @Description: get by travel detail id
// @Date: 2021-5-26 11:02 pm.
//根据 主键 查询旅游团
func (*TravelPathDetailServiceImpl) GetById(id int) *entity.TravelPathDetail {
	detail := &entity.TravelPathDetail{}
	db.First(detail, id)
	return detail
}
