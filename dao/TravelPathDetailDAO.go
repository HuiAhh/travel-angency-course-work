package dao

import (
	"travel-agency/core/dto"
	"travel-agency/entity"
	"travel-agency/entity/resultMap"
)

type TravelPathDetailDAO interface {

	SaveTravelPathDetail(obj *entity.TravelPathDetail) (int, error)

	//分页查询 根据 旅游路线 ID 查询出旅游团
	FindListTravelPathDetail(query *entity.TravelPathDetail,page int ,size int) (dto.PageResult,[]resultMap.TravelDetailWithOrderPersonCount)


}

