package entity

import (
	log "github.com/sirupsen/logrus"
	"time"
	"travel-agency/entity/base"
	"travel-agency/util"
)

type TravelPathDetail struct {
	TravelDetailId int `gorm:"primary_key;->" validate:"" json:"travelDetailId"`  //旅游团ID
	TravelPathId  *int   `json:"travelPathId"` // 路线 ID

	MaxPersonCount *int `json:"maxPersonCount"`	//多少个人
	BeginTime  *base.Time `json:"beginTime"` //预定开始时间
	EndTime   *base.Time `json:"endTime"`	//预定结束时间
	TeamName *string `json:"teamName"`			//旅游团队伍名字
	TeamDescription *string `json:"teamDescription"` //旅游团描述
	DeleteStatus   *bool  //逻辑删除
	RegisterEndTime *base.Time `json:"registerEndTime"` //报名截止时间
	ActiveStatus *bool  `json:"activeStatus"`// 是否可以报名

	GmtCreate         *time.Time `gorm:"autoCreateTime" json:"gmtCreate"`   //创建时间
	GmtModified       *time.Time `gorm:"autoUpdateTime" json:"gmtModified"` //修改时间
}

func (TravelPathDetail) TableName() string {
	return "t_travel_path_detail"
}

//重写 toString
func (cur *TravelPathDetail) String() string {
	return util.ObjectToStr(cur)
}

func (u *TravelPathDetail) BeforeSave() (err error) {
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

func (u *TravelPathDetail) BeforeUpdate() (err error) {
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


