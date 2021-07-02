package association

import "time"

type OrderDetailAssociation struct {
	//detail
	OrderDetailId          int
	TravelPathId           int
	PersonCount            *int
	ChildCount             *int
	OrderId                int
	OrderDetailDescription *string

	//order
	TravelDetailId *int
	PayStatus      *bool
	DeleteStatus   *bool
	GmtCreate      *time.Time
	GmtModified    *time.Time
	UpdateByUserid *int
}
