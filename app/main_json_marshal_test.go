package main_test

import (
	"testing"
	"travel-agency/entity"
)

func TestJsonMarshal(t *testing.T) {
	u := &entity.User{
		Id:           1,
		Username:     "hello",
		Password:     "1111",
		UserDetail:   "dddd",
		DeleteStatus: false,
	}
	stp, _ := u.StringPtr()
	t.Logf("res: %s", *stp)
}
