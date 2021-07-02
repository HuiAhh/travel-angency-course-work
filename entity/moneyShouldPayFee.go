package entity

type MoneyShouldPayFee struct {
	Id             int `gorm:"primary_key;->" validate:"" json:"id"`
	DistanceTimeDays   *int `json:"distanceTimeDays"`
	PayRatio       *float64 `json:"payRatio"`
	CancelPayRatio *float64 `json:"cancelPayRatio"`
}

func (MoneyShouldPayFee) TableName() string {
	return "t_money_should_pay_fee"
}
