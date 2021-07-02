package test

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testing"
	"travel-agency/dao/impl"
	"travel-agency/entity"
	"travel-agency/factory/dao"
	_ "travel-agency/factory/db"
	"travel-agency/factory/service"

	//_ "travel-agency/factory/dao"
)

func TestPrint(t *testing.T) {
	//w:= impl.TravelPathDetailDAOImpl{}
	//println(w)
	//printf("hello world %v")
}
func TestSaveOrUpdate11111(t *testing.T) {
	//w:= dao.GetTravelPathDetailDao()
	//s := "hello"
	//res ,_ := w.SaveTravelPathDetail(&entity.TravelPathDetail{
	//	TeamName: &s,
	//})
	//println(res)
	fmt.Println("sss")

}
func TestName111(t *testing.T) {
	println("ss")
	w := impl.TravelPathDAOImpl{}
	w2 := dao.GetTravelPathDetailDao()
	fmt.Println(w)
	fmt.Println(w2)
	s := "hello world"
	id := 14
	w2.SaveTravelPathDetail(&entity.TravelPathDetail{
		TeamName:       &s,
		TravelPathId:  &id,

		TravelDetailId: 1112,
	})
}

func Test2222(t *testing.T) {
	w := service.GetTravelPathDetailService()
	id := 10
	name := ""
	res, _ := w.QueryDetailInfoByPathId(&entity.TravelPathDetail{TravelPathId: &id,TeamName: &name},1,10)
	fmt.Println(&res)
}



func TestTravelDetailWith_CNT(t *testing.T) {
	DAO := dao .GetTravelPathDetailDao()
	id := 10;
	teamName := ""
	pageresult, list := DAO.FindListTravelPathDetail(&entity.TravelPathDetail{
		TravelPathId: &id,
		TeamName: &teamName,
	},1,10)
	fmt.Println(&pageresult)
	fmt.Println(&list[0])
}