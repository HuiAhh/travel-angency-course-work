package entity

type OrderDetail struct {
	OrderDetailId          int  `gorm:"primary_key;->" validate:""`
	PersonCount            *int `validate:"number,gt=0"`
	ChildCount             *int `validate:"number"`
	OrderId                *int
	OrderDetailDescription *string
}

func (OrderDetail) TableName() string {
	return "t_order_detail"
}
