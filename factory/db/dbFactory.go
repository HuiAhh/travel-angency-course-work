package db

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	dbUtil "travel-agency/db"
)

//var 	Db, _  (*gorm.DB,error ) =dbUtil.GetDBConnection()
var mydb *gorm.DB

func init() {
	mydb, _ = dbUtil.GetDBConnection()
	log.Println("加载 db __", mydb)

}

func GetGormDB() *gorm.DB {
	if mydb == nil {
		log.Warn("空指针异常 ，db")
	}
	return mydb
}
