package service

import "travel-agency/entity"

type UserService interface {
	//用户登录
	Login(u *entity.User) bool
	// Register(u *entity.User, rePassword string) bool
}
