package dao

import "travel-agency/entity"

type UserDao interface {
	Create(u *entity.User)
	GetAll() []entity.User
	GetById(id int) *entity.User
	Update(u *entity.User)
	DeleteById(id int)

	CheckUserNameAndPassword(u *entity.User) bool
}
