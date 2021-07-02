package entity

import "travel-agency/util"

// 用户联系方式详情
type PersonInfo struct {
//<<<<<<< HEAD
	PersonId      int `gorm:"primary_key;->" validate:"" json:"personId"`
	PersonName    string  `json:"personName"` //用户名字
	Email         string `json:"email"` //邮箱
	OrderId int             `json:"orderId"`	//订单信息
	DeleteStatus  bool      `json:"deleteStatus"`//逻辑删除
	Phone         string `json:"phone"`	//电话号码
	PersonAddress string `json:"personAddress"` //联系地址
////=======
//	PersonId     int `gorm:"primary_key;->" validate:""`
//	PersonName   *string
//	Email        *string
//	OrderId      int
//	DeleteStatus *bool
//	Phone        *string
//>>>>>>> 42f77157a29fcdcacd90fbe018a863223cfa4f27
}

type PersonInfoVO struct {
	PersonInfo
}

func (PersonInfo) TableName() string {
	return "t_person_info"
}
func (x * PersonInfo ) String() string {
	return util.ObjectToStr(x)
}



func (PersonInfoVO) TableName() string {
	return "v_person_info_by_order"
}
func (x * PersonInfoVO) String() string {
	return util.ObjectToStr(x)
}

