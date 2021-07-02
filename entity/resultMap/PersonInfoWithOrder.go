package resultMap

import "travel-agency/entity"

type PersonInfoWithOrder struct {
	OrderId int 	`json:"orderId" validate:"gt=0"` //订单ID

	PersonInfoList []entity.PersonInfo 	`json:"personInfoList"` //用户信息

}
