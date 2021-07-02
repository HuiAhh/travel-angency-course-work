package dao

import "time"

type ReportDAO interface{

	//根据日期查当天收入 [放弃存储过程]
	FindPaidOrderIncomeByDay(date *time.Time) float64
	
	FindPaidOrderDepartByDay(date *time.Time) int
}