package impl

import (
	"time"
	"travel-agency/core/dto"
	"travel-agency/entity"
	"travel-agency/factory/dao"

	log "github.com/sirupsen/logrus"
)

type PayServiceImpl struct {
}

var (
	payDAO = dao.GetPayDAO()
	reportDAO=dao.GetReportDAO()
)

func (PayServiceImpl) GetAll() []entity.MoneyShouldPayFee {
	var f []entity.MoneyShouldPayFee
	db.Find(&f)
	return f
}
func (PayServiceImpl) GetNonPaidOrders(page int, size int, keyword *string) dto.PageResult {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 5
	}
	if keyword == nil {
		empty := ""
		keyword = &empty
	}
	w := &dto.PageResult{
		RequestSize: size,
		CurrentPage: page,
	}
	payDAO.FindNonPaidOrders(w, keyword)

	return *w
}

func (PayServiceImpl) GetPaidOrders(page int, size int, keyword *string) dto.PageResult {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 5
	}
	if keyword == nil {
		empty := ""
		keyword = &empty
	}
	w := &dto.PageResult{
		RequestSize: size,
		CurrentPage: page,
	}
	payDAO.FindPaidOrders(w, keyword)

	return *w
}
func (PayServiceImpl) GetNonAndPaidOrderById(id int) dto.PageResult {
	log.Printf("find byid = %d", id)

	w := &dto.PageResult{}
	payDAO.FindNonAndPaidOrderById(w, id)
	return *w
}
func (PayServiceImpl) GetWeeklyIncome() dto.PageResult {
	type incomeData struct {
		Date   time.Time
		Income float64
	}

	// 半个月的数据
	weeklyIncomes := [15]incomeData{}
	curDate := time.Now()
	// 减一天
	dayDiff, _ := time.ParseDuration("-24h")

	//置入日期轴数据
	for i := len(weeklyIncomes)-1; i >=0; i-- {
		weeklyIncomes[i].Date = curDate

		// 查&算
		income := reportDAO.FindPaidOrderIncomeByDay(&curDate)

		weeklyIncomes[i].Income = income
		// DATEDIFF par1后, par2先 结果为正值
		// db.Where("DATEDIFF(gmt_modified,NOW()) =0").Find(&weeklyIncomes[i])

		//下一个
		curDate=curDate.Add(dayDiff)
	}
	d:=dto.PageResult{}
	d.Result= weeklyIncomes
	return d
}

// 未来半个月的开团
func (PayServiceImpl) GetWeeklyDepart() dto.PageResult {
	type incomeData struct {
		Date   time.Time
		Amount int
	}

	// 半个月的数据
	weeklyIncomes := [15]incomeData{}
	curDate := time.Now()
	// 减一天
	dayDiff, _ := time.ParseDuration("24h")

	//置入日期轴数据
	for i,length := 0,len(weeklyIncomes); i <length; i++ {
		weeklyIncomes[i].Date = curDate

		// 查&算
		amount := reportDAO.FindPaidOrderDepartByDay(&curDate)

		weeklyIncomes[i].Amount = amount
		// DATEDIFF par1后, par2先 结果为正值
		// db.Where("DATEDIFF(gmt_modified,NOW()) =0").Find(&weeklyIncomes[i])

		//下一个
		curDate=curDate.Add(dayDiff)
	}
	d:=dto.PageResult{}
	d.Result= weeklyIncomes
	return d
}