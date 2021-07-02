package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
	"testing"
)

func TestGetDBConnection(t *testing.T) {
	tests := []struct {
		name    string
		want    *gorm.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDBConnection()
			fmt.Println(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDBConnection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDBConnection() got = %v, want %v", got, tt.want)
			}
		})
	}
}
