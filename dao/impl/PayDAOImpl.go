package impl

import (
	"fmt"
	"log"
	"time"
	"travel-agency/core/dto"
	"travel-agency/entity"
	"travel-agency/util"
)

type PayDAOImpl struct {
}

type recordsCount struct{
	// 结果行数
	ResCount int
}

// 未支付订单 带order id模糊查询
func (PayDAOImpl) FindNonPaidOrders(result *dto.PageResult, keyword *string) {
	if result == nil {
		return
	}
	//1+n查询
	//1.order
	var nonAndPaidOrders []nonAndPaidOrder
	db.Raw(
		"select order_id, travel_path_id, travel_detail_id, pay_status, gmt_create "+
			"from t_order where delete_status=false and pay_status=false and order_id like ?", fmt.Sprintf("%%%s%%", *keyword)).
		Offset((result.CurrentPage - 1) * result.RequestSize).
		Limit(result.RequestSize).
		Order("gmt_create desc").
		Scan(&nonAndPaidOrders)

	// 分页

	// db.Model(&entity.Order{}).Count(&result.TotalCount)
	rCount:=&recordsCount{}
	db.Raw(
		"select count(*) as res_count "+
			"from t_order where delete_status=false and pay_status=false and order_id like ?", fmt.Sprintf("%%%s%%", *keyword)).
		Scan(rCount)
	result.TotalCount=rCount.ResCount
	
	//总页数
	result.TotalPage = ((result.TotalCount - 1) / result.RequestSize) + 1

	log.Printf("from db: %v", util.ObjectToStr(nonAndPaidOrders))
	length := len(nonAndPaidOrders)
	for i := 0; i < length; i++ {
		//
		getOtherFields1(&nonAndPaidOrders[i])

		// 获取费率<-获取日期差值
		diff := util.GetDayDiff(nonAndPaidOrders[i].GmtCreate, nonAndPaidOrders[i].BeginTime)
		if diff >= 0 {
			// 匹配区间
			dayLevel := 0
			if diff < 30 {
				dayLevel = 10
			} else if diff < 60 {
				dayLevel = 30
			} else {
				dayLevel = 60
			}

			payFee := &entity.MoneyShouldPayFee{DistanceTimeDays: &dayLevel}
			db.Where(payFee).First(payFee)

			receivableRatio := 1.0 - *payFee.PayRatio

			//5.dynatic account receivable 应收账款[余款]

			var childCost float64 = float64(nonAndPaidOrders[i].ChildCount) * (nonAndPaidOrders[i].ChildShouldPay)
			var adultCost float64 = float64(nonAndPaidOrders[i].PersonCount-nonAndPaidOrders[i].ChildCount) * (nonAndPaidOrders[i].PerShouldPay)
			var receivable float64 = (childCost + adultCost) * receivableRatio
			nonAndPaidOrders[i].AccountReceivable = receivable
		}
	}

	result.Result = nonAndPaidOrders
}

//已支付订单 带order id模糊查询
func (PayDAOImpl) FindPaidOrders(result *dto.PageResult, keyword *string) {
	if result == nil {
		return
	}
	//1+n查询
	//1.order
	var nonAndPaidOrders []nonAndPaidOrder
	db.Raw(
		"select order_id, travel_path_id, travel_detail_id, pay_status, gmt_create "+
			"from t_order where delete_status=false and pay_status=true and order_id like ?", fmt.Sprintf("%%%s%%", *keyword)).
		Offset((result.CurrentPage - 1) * result.RequestSize).
		Limit(result.RequestSize).
		Order("gmt_create desc").
		Scan(&nonAndPaidOrders)

	// 分页

	// db.Model(&entity.Order{}).Count(&result.TotalCount)

	rCount:=&recordsCount{}
	db.Raw(
		"select count(*) as res_count "+
			"from t_order where delete_status=false and pay_status=true and order_id like ?", fmt.Sprintf("%%%s%%", *keyword)).
		Scan(rCount)
	result.TotalCount=rCount.ResCount

	//总页数
	result.TotalPage = ((result.TotalCount - 1) / result.RequestSize) + 1

	log.Printf("from db: %v", util.ObjectToStr(nonAndPaidOrders))
	length := len(nonAndPaidOrders)
	for i := 0; i < length; i++ {
		getOtherFields1(&nonAndPaidOrders[i])

		//5.dynatic account receivable 应收账款[总]

		var childCost float64 = float64(nonAndPaidOrders[i].ChildCount) * (nonAndPaidOrders[i].ChildShouldPay)
		var adultCost float64 = float64(nonAndPaidOrders[i].PersonCount-nonAndPaidOrders[i].ChildCount) * (nonAndPaidOrders[i].PerShouldPay)
		var receivable float64 = (childCost + adultCost)
		nonAndPaidOrders[i].AccountReceivable = receivable
	}

	result.Result = nonAndPaidOrders
}

//根据order id 查
func (PayDAOImpl) FindNonAndPaidOrderById(result *dto.PageResult, id int) {
	if result == nil {
		return
	}
	//1+n查询
	//1.order
	nonAndPaid := &nonAndPaidOrder{}
	db.Raw(
		"select order_id, travel_path_id, travel_detail_id, pay_status, gmt_create "+
			"from t_order where delete_status=false and order_id = ?", id).
		Scan(&nonAndPaid)
	log.Printf("from db: %v", util.ObjectToStr(nonAndPaid))

	//2....
	getOtherFields1(nonAndPaid)

	//5.dynatic account receivable 应收账款[总]

	var childCost float64 = float64(nonAndPaid.ChildCount) * (nonAndPaid.ChildShouldPay)
	var adultCost float64 = float64(nonAndPaid.PersonCount-nonAndPaid.ChildCount) * (nonAndPaid.PerShouldPay)
	var receivable float64 = (childCost + adultCost)
	nonAndPaid.AccountReceivable = receivable

	result.Result = nonAndPaid

}

func getOtherFields1(napo *nonAndPaidOrder) {
	//2.detail
	type tCount struct {
		//order detail
		PersonCount   int
		ChildCount    int
		OrderDetailId int
	}
	tempCount := &tCount{}
	db.Raw("select person_count, child_count, order_detail_id "+
		"from t_order o left join t_order_detail od on o.order_id=od.order_id "+
		"where o.order_id=?", napo.OrderId).
		Scan(tempCount)
	napo.PersonCount = tempCount.PersonCount
	napo.ChildCount = tempCount.ChildCount
	napo.OrderDetailId = tempCount.OrderDetailId

	//3.travel path
	type tPath struct {
		//travel path name...
		TravelPathName *string
		PerShouldPay   float64
		ChildShouldPay float64
	}
	tempPath := &tPath{}
	db.Raw("select travel_path_name, per_should_pay, child_should_pay "+
		"from t_order o left join t_travel_path tp on o.travel_path_id=tp.travel_path_id "+
		"where o.travel_path_id=?", napo.TravelPathId).
		Scan(tempPath)
	napo.TravelPathName = tempPath.TravelPathName
	napo.PerShouldPay = tempPath.PerShouldPay
	napo.ChildShouldPay = tempPath.ChildShouldPay

	//4.travel path detail
	type tTime struct {
		BeginTime       *time.Time
		TeamName        *string
		TeamDescription *string
	}
	tempTime := &tTime{}
	db.Raw("select begin_time, team_name, team_description "+
		"from t_order o left join t_travel_path_detail tpd on o.travel_detail_id=tpd.travel_detail_id "+
		"where o.travel_detail_id=?", napo.TravelDetailId).
		Scan(tempTime)
	napo.BeginTime = tempTime.BeginTime
	napo.TeamName = tempTime.TeamName
	napo.TeamDescription = tempTime.TeamDescription
}

type nonAndPaidOrder struct {
	//order
	OrderId        int
	OrderDetailId  int
	TravelPathId   int
	TravelDetailId int
	PayStatus      *bool
	GmtCreate      *time.Time

	//order detail
	PersonCount int
	ChildCount  int

	//travel path name...
	TravelPathName *string
	PerShouldPay   float64
	ChildShouldPay float64

	//travel path detail 旅游团
	BeginTime       *time.Time
	TeamName        *string
	TeamDescription *string

	AccountReceivable float64 //应收账款 动态计算出来


}
