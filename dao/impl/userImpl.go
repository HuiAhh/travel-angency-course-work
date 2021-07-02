package impl

import (
	"travel-agency/entity"

	"github.com/jinzhu/gorm"
)

type UserDao struct {
	DB *gorm.DB
}

// &User{} 还是会向数据库传0值，而不是null
func (d *UserDao) Create(u *entity.User) {
	d.DB.Create(u)
}

func (d *UserDao) GetAll() []entity.User {
	users := []entity.User{}
	d.DB.Find(&users)
	return users
}

// 查出id为0意味着是空值，这样省去判nil操作
func (d *UserDao) GetById(id int) *entity.User {
	user := entity.User{}
	d.DB.First(&user, id)
	return &user
}

func (d *UserDao) Update(u *entity.User) {
	d.DB.Save(u)
}

func (d *UserDao) DeleteById(id int) {
	d.DB.Delete(&entity.User{}, id)
}

//自定义crud

func (d *UserDao) CheckUserNameAndPassword(u *entity.User) bool {
	result := d.DB.Where("username=? and password=?", u.Username, u.Password).First(&entity.User{})
	return result.RowsAffected > 0
}
