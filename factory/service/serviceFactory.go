package service

import (
	log "github.com/sirupsen/logrus"
	"travel-agency/service"
	"travel-agency/service/impl"
)

//user service 实现类
var (
	user_service       impl.UserServiceImpl
	travelPath_service impl.TravelServiceImpl
	order_service impl.OrderServiceImpl
	pay_service impl.PayServiceImpl

	travelPathDetail_service impl.TravelPathDetailServiceImpl
	personInfo_service impl.PersonInfoServiceImpl
)

func init() {
	//user_service = impl.UserServiceImpl{}
	//travelPath_service = impl.TravelServiceImpl{}
	//travelPathDetail_service = impl.TravelPathDetailServiceImpl{}
	//log.SetFlags(flag.)
	log.Printf("加载工厂包")
}

// func init() {

// 	//log.SetFlags(flag.)
// 	log.Printf("加载工厂包")
// }

//工厂模式 获取 具体实现
func GetUserService() service.UserService {
	return &user_service
}
func GetTravelPathService() service.TravelService {
	return &travelPath_service
}
func GetOrderService() service.OrderService{
	return &order_service
}
func GetPayService() service.PayService{
	return &pay_service
}

func GetTravelPathDetailService() service.TravelDetailService {

	return &travelPathDetail_service
}
func GetPersonInfoService() service.PersonInfoService{
	return &personInfo_service
}