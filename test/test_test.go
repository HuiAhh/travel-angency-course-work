package test

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testing"
	"travel-agency/entity"
	"travel-agency/factory/dao"
	dbfactory "travel-agency/factory/db"
	//_ "travel-agency/factory/dao"
)

func Test_9(t *testing.T) {
	db := dbfactory.GetGormDB()
	fmt.Println(db)
}

func TestSaveOrUpdate111(t *testing.T) {
	w:= dao.GetTravelPathDetailDao()
	s := "hello"
	res ,_ := w.SaveTravelPathDetail(&entity.TravelPathDetail{
		TeamName: &s,
	})
	println(res)

}

////查询
//func Test_findlist(t *testing.T) {
//	//pathdao.
//	s := dto.PageResult{
//		CurrentPage: 1,
//		RequestSize: 5,
//	}
//	var pathdao = dao.GetTravelPathDao()
//	//list := pathdao.FindList(&entity.TravelPath{ActiveStatus: true}, &s)
//	fmt.Println(s)
//	fmt.Println("----------n")
//	for i, v := range list {
//		fmt.Println(i)
//		fmt.Println(v)
//	}
//}

//查询
func Test_findlist22(t *testing.T) {
	var db = dbfactory.GetGormDB()
	fmt.Println(db)
	//pathdao.
	var res []entity.TravelPath
	db.
		Find(&res)
	fmt.Println(res)
	//Limit(result.RequestSize).
	//Offset((result.CurrentPage-1) *result.RequestSize)
}
