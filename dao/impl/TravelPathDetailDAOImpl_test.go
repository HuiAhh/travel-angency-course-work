package impl

import (
	"fmt"
	"testing"
	"travel-agency/core/dto"
	"travel-agency/entity"
)

func TestAPP111(t *testing.T) {
	w := TravelPathDetailDAOImpl{}
	id := 14
	detail, res := w.FindListTravelPathDetail(&entity.TravelPathDetail{
		TravelPathId: &id,
	}, 1, 10)
	fmt.Printf("res = %v,x = %v\n", detail,res)
	detail.TotalPage = -1;
	detail.CalcTotalPage()
	fmt.Println(detail.TotalPage)
}

func TestName111(t *testing.T) {
	result := dto.PageResult{}
	result.RequestSize = 10
	result.CurrentPage = 1
	result.TotalCount = 100
	result.CalcTotalPage()
	//println(result)
	fmt.Println(result)
}