package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

func GetDBConnection() (*gorm.DB, error) {
	// db
	//GMT +8 ,设置中国时区 改本地
	dbLink := "root001:root@tcp(localhost:3309)/webapp001?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", dbLink)
	if err != nil {
		log.Warn("连接异常  ///xxx", db, err)
		return nil, err
	}
	db.LogMode(true)
	return db, nil
}
