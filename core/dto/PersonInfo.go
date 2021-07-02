package dto

import "travel-agency/entity"

type PersonInfoDTO struct {
	//订单详情
	OrderDetailId int  `json:"orderDetailId" validate:"gt:0"`
	PersonInfo []entity.PersonInfo `json:"personInfo"`



}

