package dao

import (
	"travel-agency/core/dto"
	"travel-agency/entity"
)

type TravelPathDAO interface {
	SaveOrUpdate(obj *entity.TravelPath) (int, error)

	FindList(query *entity.TravelPath, result *dto.PageResult)
}
