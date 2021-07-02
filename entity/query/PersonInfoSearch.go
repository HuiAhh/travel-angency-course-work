package query

import "travel-agency/util"

type PersonInfoSearch struct {
	TravelDetailId int  `form:"travelDetailId"`
	TravelPathId int `form:"travelPathId"`
	PersonName string `form:"personName"`
	//PersonId *int
	OrderId int `form:"orderId"`
}

func (PersonInfoSearch)TableName() string {
	return "v_person_info_by_order"
}
func (c *PersonInfoSearch) String() string {
	return util.ObjectToStr(c)
}
//db.Model(query).Find()