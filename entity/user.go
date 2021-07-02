package entity

import (
	"encoding/json"
	"travel-agency/util"
)

// model, dao, service
// 结构体对应默认gorm的命名法：字段UserInfo -> user_info  解析为小写下划线
//						结构体名User -> users      解析为加复数
type User struct {
	// 箭头方向参阅 https://gorm.io/docs/models.html
	// 箭头可能只有在save()下有用
	Id           int    `gorm:"primary_key;->" validate:""`
	Username     string `validate:"required,email"`
	Password     string `validate:"is-password"`
	UserDetail   string
	DeleteStatus bool `gorm:"->;<-:update"`
	Salt string //盐密码
}

func (User) TableName() string {
	return "t_user"
}

func (u User) StringPtr() (*string, error) {
	res, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	ret := string(res)
	return &ret, nil
}
func (u *User) String() string {
	return util.ObjectToStr(u)
}