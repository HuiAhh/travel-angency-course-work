package dao

import (
	daoImpl "travel-agency/dao/impl"

	log "github.com/sirupsen/logrus"

	//"log"
	//"travel-agency/dao"
	//dbUtil "travel-agency/db"
	//
	dbFactory "travel-agency/factory/db"
	//// serviceImpl "travel-agency/service/impl"

	"travel-agency/dao"

	_ "github.com/go-playground/validator"
	_ "github.com/jinzhu/gorm"
)

//user service 实现类
var (
	userDao       dao.UserDao
	travelPathDAO dao.TravelPathDAO
	orderDAO      dao.OrderDAO
	payDAO dao.PayDAO
	travelPathDetailDAO dao.TravelPathDetailDAO
	reportDAO dao.ReportDAO


	//非初始化
	personInfoDAO daoImpl.PersonInfoDAOImpl
)

func init() {

	//db = dbf.GetDB()

	//dao
	userDao = &daoImpl.UserDao{DB: dbFactory.GetGormDB()}
	travelPathDAO = &daoImpl.TravelPathDAOImpl{}
	orderDAO = &daoImpl.OrderDAOImpl{}
	payDAO=&daoImpl.PayDAOImpl{}

	travelPathDetailDAO = &daoImpl.TravelPathDetailDAOImpl{}
	reportDAO=&daoImpl.ReportDAOImpl{}
	// //service
	// userService = &serviceImpl.UserServiceImpl{}

	//validate

}

//dao
func GetUserDao() dao.UserDao {
	return userDao
}

func GetTravelPathDao() dao.TravelPathDAO {
	if travelPathDAO == nil {
		log.Fatal("工厂方法空指针")
	}
	return travelPathDAO
}

func GetOrderDAO() dao.OrderDAO {
	if orderDAO == nil {
		log.Fatal("工厂方法空指针")
	}
	return orderDAO
}
func GetPayDAO() dao.PayDAO {
	if payDAO == nil {
		log.Fatal("工厂方法空指针")
	}
	return payDAO
}

func GetTravelPathDetailDao() dao.TravelPathDetailDAO {
	if travelPathDetailDAO == nil{
		log.Warn("空指针异常 ，travelPathDetailDAO")
	}
	return travelPathDetailDAO
}

func GetReportDAO() dao.ReportDAO {
	if reportDAO == nil {
		log.Fatal("工厂方法空指针")
	}
	return reportDAO
}

func GetPersonInfoDAO() dao.PersonInfoDAO {
	return &personInfoDAO
}