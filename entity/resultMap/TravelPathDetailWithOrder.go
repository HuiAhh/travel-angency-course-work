package resultMap

import (
	"travel-agency/entity"
	"travel-agency/util"
)

type TravelDetailWithOrderPersonCount struct {

	entity.TravelPathDetail // 旅游团信息

	CurrentPersonCount int 	//当前报名人数

}


func (c * TravelDetailWithOrderPersonCount) String() string {
	return util.ObjectToStr(c)
}

//func a() {
//	w := TravelDetailWithOrder{}
//	id := w.TravelDetailId
//	fmt.Println(w.currentPersonCount)
//	fmt.Println(id)
//}