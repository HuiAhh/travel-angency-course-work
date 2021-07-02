package entity

import (
	log "github.com/sirupsen/logrus"
	"time"
	"travel-agency/util"
)

type TravelPath struct {
	TravelPathId   int     `gorm:"primary_key;->"   json:"travelPathId,int" ` //主键 ID
	TravelPathName *string `validate:"" json:"travelPathName"`                //旅游路径名字

	TravelPathContent *string `validate:""  json:"travelPathContent"` //旅游路径内容

	DeleteStatus      *bool      `gorm:"->;<-:update" validate:""  json:"deleteStatus"`
	GmtCreate         *time.Time `gorm:"autoCreateTime" json:"gmtCreate"`   //创建时间
	GmtModified       *time.Time `gorm:"autoUpdateTime" json:"gmtModified"` //修改时间
	ActiveStatus      *bool      `json:"activeStatus"`                      //是否可以下订单， 1表示 active，可以下订单， 0 表示不可以下订单
	UpdateBy          *string    //被谁更新过
	TravelDuration    *string    `json:"travelDuration"`                              //持续时间,单位天
	TravelDescription *string    `validate:"required,min=1" json:"travelDescription"` // 内容概要
	BeginArea         *string    `validate:"min=1" json:"beginArea"`                  //出发地点
	DestinationArea   *string    `json:"destinationArea"`                             //目的地点
	PerShouldPay      *float64   `validate:"number,gt=0" json:"perShouldPay"`         //每个成年人 应该 付款多少钱
	ChildShouldPay    *float64   `validate:"number,gt=0"  json:"childShouldPay" `     //每个孩子应该付款多少钱
}

//TODO:
//神坑， json 的 PerShouldPay 字符串 "44.11" 转不了 float，还要 在 tag 标注才能转
//https://blog.csdn.net/iamlihongwei/article/details/78842612

func (TravelPath) TableName() string {
	return "t_travel_path"
}
func (u *TravelPath) String() string {
	return util.ObjectToStr(u)
}

//func (TravelPath) String() string{
//	return fmt.Sprintf("travelPathid")
//}

// func (t TravelPath) String() string {

// }

func (u *TravelPath) BeforeSave() (err error) {
	if u != nil {
		if u.GmtCreate == nil {
			x := time.Now()
			//创建时间
			u.GmtCreate = &x
		}
		if u.GmtModified == nil {
			x := time.Now()
			//创建时间
			u.GmtModified = &x
		}
		if u.DeleteStatus == nil {
			w := false
			u.DeleteStatus = &w
		}
	} else {
		log.Warn("空指针异常 travelPath")
	}
	return
}

func (u *TravelPath) BeforeUpdate() (err error) {
	if u != nil {

		if u.GmtModified == nil {
			x := time.Now()
			//创建时间
			u.GmtModified = &x
		}
	} else {
		log.Warn("空指针异常 travelPath")
	}
	return
}
