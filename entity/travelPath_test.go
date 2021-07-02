package entity

import (
	_ "fmt"
	log "github.com/sirupsen/logrus"
	"testing"

	_ "travel-agency/factory/dao"
)

func Test_TOsTRING(t *testing.T)  {
	w := "sss"
	u := TravelPath{
		TravelPathId: 1,
		TravelPathContent:  &w,
	}
	log.Println(u)
	log.Println(&u)
	//log.Println(u.ToString())
	//log.Println(util.ObjectToStr(u))
}


//func TestMoneyShouldPayFee_TableName(t *testing.T) {
//	type fields struct {
//		Id             int
//		DistanceTime   string
//		PayRatio       string
//		CancelPayRatio string
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			mo := MoneyShouldPayFee{
//				Id:             tt.fields.Id,
//				DistanceTime:   tt.fields.DistanceTime,
//				PayRatio:       tt.fields.PayRatio,
//				CancelPayRatio: tt.fields.CancelPayRatio,
//			}
//			if got := mo.TableName(); got != tt.want {
//				t.Errorf("TableName() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestOrderDetail_TableName(t *testing.T) {
//	type fields struct {
//		OrderDetailId          int
//		PersonCount            int
//		OrderId                int
//		OrderDetailDescription string
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			or := OrderDetail{
//				OrderDetailId:          tt.fields.OrderDetailId,
//				PersonCount:            tt.fields.PersonCount,
//				OrderId:                tt.fields.OrderId,
//				OrderDetailDescription: tt.fields.OrderDetailDescription,
//			}
//			if got := or.TableName(); got != tt.want {
//				t.Errorf("TableName() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestOrder_TableName(t *testing.T) {
//	type fields struct {
//		OrderId        int
//		TravelPathId   int
//		PayStatus      int
//		DeleteStatus   int
//		GmtCreate      time.Time
//		GmtModified    time.Time
//		UpdateByUserid int
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			or := Order{
//				OrderId:        tt.fields.OrderId,
//				TravelPathId:   tt.fields.TravelPathId,
//				PayStatus:      tt.fields.PayStatus,
//				DeleteStatus:   tt.fields.DeleteStatus,
//				GmtCreate:      tt.fields.GmtCreate,
//				GmtModified:    tt.fields.GmtModified,
//				UpdateByUserid: tt.fields.UpdateByUserid,
//			}
//			if got := or.TableName(); got != tt.want {
//				t.Errorf("TableName() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestPersonInfo_TableName(t *testing.T) {
//	type fields struct {
//		PersonId      int
//		PersonName    string
//		Email         string
//		OrderDetailId int
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			pe := PersonInfo{
//				PersonId:      tt.fields.PersonId,
//				PersonName:    tt.fields.PersonName,
//				Email:         tt.fields.Email,
//				OrderDetailId: tt.fields.OrderDetailId,
//			}
//			if got := pe.TableName(); got != tt.want {
//				t.Errorf("TableName() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestTravelPathDetail_TableName(t *testing.T) {
//	type fields struct {
//		TravelDetailId int
//		DetailTimes    string
//		TravelPathId   int
//		MaxPersonCount int
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tr := TravelPathDetail{
//				TravelDetailId: tt.fields.TravelDetailId,
//				DetailTimes:    tt.fields.DetailTimes,
//				TravelPathId:   tt.fields.TravelPathId,
//				MaxPersonCount: tt.fields.MaxPersonCount,
//			}
//			if got := tr.TableName(); got != tt.want {
//				t.Errorf("TableName() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestTravelPath_TableName(t *testing.T) {
//	type fields struct {
//		TravelPathId      int
//		TravelPathName    string
//		TravelPathContent string
//		DeleteStatus      bool
//		GmtCreate         *time.Time
//		GmtModified       *time.Time
//		ActiveStatus      bool
//		UpdateBy          string
//		TravelDuration    string
//		TravelDescription string
//		BeginArea         string
//		DestinationArea   string
//		PerShouldPay      float64
//		ChildShouldPay    float64
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tr := TravelPath{
//				TravelPathId:      tt.fields.TravelPathId,
//				TravelPathName:    tt.fields.TravelPathName,
//				TravelPathContent: tt.fields.TravelPathContent,
//				DeleteStatus:      tt.fields.DeleteStatus,
//				GmtCreate:         tt.fields.GmtCreate,
//				GmtModified:       tt.fields.GmtModified,
//				ActiveStatus:      tt.fields.ActiveStatus,
//				UpdateBy:          tt.fields.UpdateBy,
//				TravelDuration:    tt.fields.TravelDuration,
//				TravelDescription: tt.fields.TravelDescription,
//				BeginArea:         tt.fields.BeginArea,
//				DestinationArea:   tt.fields.DestinationArea,
//				PerShouldPay:      tt.fields.PerShouldPay,
//				ChildShouldPay:    tt.fields.ChildShouldPay,
//			}
//			if got := tr.TableName(); got != tt.want {
//				t.Errorf("TableName() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestUser_StringPtr(t *testing.T) {
//	type fields struct {
//		Id           int
//		Username     string
//		Password     string
//		UserDetail   string
//		DeleteStatus bool
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		want    *string
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			u := User{
//				Id:           tt.fields.Id,
//				Username:     tt.fields.Username,
//				Password:     tt.fields.Password,
//				UserDetail:   tt.fields.UserDetail,
//				DeleteStatus: tt.fields.DeleteStatus,
//			}
//			got, err := u.StringPtr()
//			if (err != nil) != tt.wantErr {
//				t.Errorf("StringPtr() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("StringPtr() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestUser_TableName(t *testing.T) {
//	type fields struct {
//		Id           int
//		Username     string
//		Password     string
//		UserDetail   string
//		DeleteStatus bool
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			us := User{
//				Id:           tt.fields.Id,
//				Username:     tt.fields.Username,
//				Password:     tt.fields.Password,
//				UserDetail:   tt.fields.UserDetail,
//				DeleteStatus: tt.fields.DeleteStatus,
//			}
//			if got := us.TableName(); got != tt.want {
//				t.Errorf("TableName() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
