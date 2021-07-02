package impl

import (
	"time"
	"travel-agency/util"
)

type ReportDAOImpl struct{}

//根据日期查当天收入 [放弃存储过程]
func (ReportDAOImpl) FindPaidOrderIncomeByDay(date *time.Time) float64 {
	var totalIncome float64 = 0
	var personTotal float64
	var childTotal float64

	//1.查当天支付记录、+余款[]: gmt_modified=[day_date], pay_status=true, delete_status=false
	paidOrders := []record{}
	db.Raw("select order_id, travel_path_id, travel_detail_id, gmt_create, gmt_modified "+
		"from t_order "+
		"where DATEDIFF(gmt_modified, ?)=0 and pay_status=true and delete_status=false", date).
		Scan(&paidOrders)
	for i, length := 0, len(paidOrders); i < length; i++ {
		otherField := getOtherFields(&paidOrders[i], date, false)
		personTotal = (otherField.PerShouldPay) * (float64(otherField.PersonCount))
		childTotal = (otherField.ChildShouldPay) * (float64(otherField.ChildCount))
		totalIncome += (personTotal + childTotal) * (1 - otherField.PayRatio)
	}

	refundOrders := []record{}
	//2.查当天撤单记录、-退款[]: gmt_modified=[day_date], pay_status=true, delete_status=true
	db.Raw("select order_id, travel_path_id, travel_detail_id, gmt_create, gmt_modified "+
		"from t_order "+
		"where DATEDIFF(gmt_modified, ?)=0 and pay_status=true and delete_status=true", date).
		Scan(&refundOrders)
	for i, length := 0, len(refundOrders); i < length; i++ {
		otherField := getOtherFields(&refundOrders[i], date, true)
		personTotal = (otherField.PerShouldPay) * (float64(otherField.PersonCount))
		childTotal = (otherField.ChildShouldPay) * (float64(otherField.ChildCount))
		totalIncome += (personTotal + childTotal) * (1 - otherField.CancelPayRatio)
	}

	applyOrders := []record{}
	//3.查当天申请记录，+定金[]: gmt_create=[day_date], pay_status=false, delete_status=false
	db.Raw("select order_id, travel_path_id, travel_detail_id, gmt_create, gmt_modified "+
		"from t_order "+
		"where DATEDIFF(gmt_create, ?)=0 and pay_status=false and delete_status=false", date).
		Scan(&applyOrders)
	for i, length := 0, len(applyOrders); i < length; i++ {
		otherField := getOtherFields(&applyOrders[i], date, false)
		personTotal = (otherField.PerShouldPay) * (float64(otherField.PersonCount))
		childTotal = (otherField.ChildShouldPay) * (float64(otherField.ChildCount))
		totalIncome += (personTotal + childTotal) * (otherField.PayRatio)
	}

	return totalIncome
}

func getOtherFields(elem *record, curDate *time.Time, refund bool) *otherFields {
	fields := otherFields{}
	//1.查客单价 1.2.3.
	db.Raw("select per_should_pay, child_should_pay "+
		"from t_travel_path "+
		"where travel_path_id=?", elem.TravelPathId).
		Scan(&fields)

	//2.查人数 1.2.3. 连表查
	db.Raw("select person_count, child_count "+
		"from t_order o left join t_order_detail od "+
		"on o.order_id=od.order_id "+
		"where o.order_id=?", elem.OrderId).
		Scan(&fields)

	//3.查出发日 1.2.3.
	db.Raw("select begin_time "+
		"from t_travel_path_detail "+
		"where travel_detail_id=?", elem.TravelDetailId).
		Scan(&fields)

	var daysDiff int
	if refund {
		//4.当天距离出发日、算退款：cancel_pay_ratio 2.
		daysDiff = util.GetDayDiff(curDate, fields.BeginTime)
	} else {
		//5.出发日距离订单创建时间、算定金、余款：pay_ratio 1.3.
		daysDiff = util.GetDayDiff(elem.GmtCreate, fields.BeginTime)
	}
	//6.算天数差范围、匹配区间 1.2.3.
	var dayLevel int = 0
	if refund {
		if daysDiff < 30 {
			dayLevel = 10
		} else if daysDiff < 60 {
			dayLevel = 30
		} else {
			dayLevel = 60
		}
	} else {
		if daysDiff < 1 {
			dayLevel = 0
		} else if daysDiff < 10 {
			dayLevel = 1
		} else if daysDiff < 30 {
			dayLevel = 10
		} else {
			dayLevel = 30
		}
	}
	//7.退款手续费、定金费率 合并2个语句 1.2.3.
	db.Raw("select cancel_pay_ratio, pay_ratio "+
		"from t_money_should_pay_fee "+
		"where distance_time_days=?", dayLevel).
		Scan(&fields)
	return &fields
}

func (ReportDAOImpl) FindPaidOrderDepartByDay(date *time.Time) int {
	// 1.where begin_time={date} in t_travel_path_detail -> travel_detail_id -> t_order
	// right join 以order表为主
	type tRet struct {
		Ret int
	}
	ret := tRet{}
	db.Raw("select count(*) as ret "+
		"from t_travel_path_detail tpd right join t_order o on tpd.travel_detail_id=o.travel_detail_id "+
		"where o.delete_status=false and o.pay_status=true and DATEDIFF(tpd.begin_time,?)=0", date).
		Scan(&ret)
	return ret.Ret
}

type record struct {
	// order
	OrderId        int
	TravelPathId   int
	TravelDetailId int
	GmtCreate      *time.Time
	GmtModified    *time.Time
}

type otherFields struct {
	// path
	PerShouldPay   float64
	ChildShouldPay float64

	// order detail
	PersonCount int
	ChildCount  int

	// path detail
	BeginTime *time.Time

	// pay fee
	PayRatio       float64
	CancelPayRatio float64
}
