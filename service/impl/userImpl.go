package impl

import (
	"travel-agency/dao"
	"travel-agency/entity"
	fd "travel-agency/factory/dao"
)

type UserServiceImpl struct {
}

var (
	userDao dao.UserDao = fd.GetUserDao()
)

func (s *UserServiceImpl) Login(u *entity.User) bool {
	return userDao.CheckUserNameAndPassword(u)
}

// func (s *UserServiceImpl) Register(u *entity.User, rePassword string) bool {

// }
