package main

import (
	"github.com/jinzhu/gorm"
	util2 "github.com/lyr-2000/goUtil/util"
	"testing"
	"travel-agency/entity"
	dbfactory "travel-agency/factory/db"
)

var db *gorm.DB = dbfactory.GetGormDB()
func TestChangePWD001(t *testing.T) {
	salt :=util2.GetRandomString(4)
	db.Create(&entity.User{
		Username: "123",
		Password: util2.Md5([]byte("123" + salt)),
		Salt: salt,
	})

}
