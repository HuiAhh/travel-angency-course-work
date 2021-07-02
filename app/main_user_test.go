package main_test

import (
	"fmt"
	"testing"
	"time"
)

// func init() {
// 	d = dao.NewUserDao()
// }

//var d *dao.UserDao

// func TestCRUDDDD(t *testing.T) {
// 	t.Log("step 1 create:")
// 	u := &entity.User{
// 		Username:     "huiahh",
// 		Password:     "123456",
// 		UserDetail:   "fxxk",
// 		DeleteStatus: time.Now(),
// 	}
// 	d.Create(u)
// 	t.Log("step 2 query all:")

// 	users := d.GetAll()
// 	for _, i := range users {
// 		t.Log(i.String())
// 	}

// 	t.Log("step 2 query find id=:")

// 	u1 := d.GetById(6)
// 	t.Log(u1.String())

// 	t.Log("step 3 update id=6:")

// 	u1.Username = "huiahhhhha"
// 	d.Update(u1)

// 	u2 := d.GetById(6)
// 	t.Log(u2.String())

// 	t.Log("step 4 delete id=6:")

// 	d.DeleteById(6)

// 	u3 := d.GetById(6)
// 	t.Log(u3.String())
// }

func Test_sss(t *testing.T) {
	fmt.Printf("result = %s", time.Now().Format("2006-01-02 15"))
}
