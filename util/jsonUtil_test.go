package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
	"travel-agency/core/enums"
)

func TestEnums(t *testing.T) {
	var x enums.ResultStatus = enums.JumpPage
	fmt.Println(x)
}

func TestIsNotBlank(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotBlank(tt.args.v); got != tt.want {
				t.Errorf("IsNotBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRenderJson(t *testing.T) {
	type args struct {
		c   *gin.Context
		obj interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestReplaceHtml(t *testing.T) {
	type args struct {
		html string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceHtml(tt.args.html); got != tt.want {
				t.Errorf("ReplaceHtml() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrToJson(t *testing.T) {
	type args struct {
		c   *gin.Context
		obj interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StrToJson(tt.args.c, tt.args.obj); (err != nil) != tt.wantErr {
				t.Errorf("StrToJson() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestToInt(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToInt(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
