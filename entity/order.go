package entity

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

type Order struct {
	OrderId        int        `gorm:"primary_key" json:"orderId,int"`
	TravelPathId   int        `json:"travelPathId" validate:"number"`
	TravelDetailId int        `json:"travelDetailId" validate:"number"` //或许要关联travel_path_detail 的主键 为了获取开团时间
	PayStatus      *bool      `json:"payStatus"`
	DeleteStatus   *bool      `json:"deleteStatus"`
	GmtCreate      *time.Time `json:"gmtCreate"`
	GmtModified    *time.Time `json:"gmtModified"` //作为支付时间
	UpdateByUserid *int       `json:"updateByUserid"`
}

func (Order) TableName() string {
	return "t_order"
}

func (u *Order) String() string {
	return fmt.Sprintf("orderId = %v, travelPathId = %v, payStatus = %v, deleteStatus = %v",
		u.OrderId, u.TravelPathId, *u.PayStatus, *u.DeleteStatus)
}

func (u *Order) BeforeSave() (err error) {
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
		log.Warn("空指针异常 order")
	}
	return
}

func (u *Order) BeforeUpdate() (err error) {
	if u != nil {

		if u.GmtModified == nil {
			x := time.Now()
			//创建时间
			u.GmtModified = &x
		}
	} else {
		log.Warn("空指针异常 order")
	}
	return
}
