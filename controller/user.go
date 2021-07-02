package controller

import (
	"net/http"

	"travel-agency/entity"

	ff "travel-agency/factory/service"
	"travel-agency/service"

	"github.com/gin-gonic/gin"
)

var (
	userService service.UserService = ff.GetUserService()
)

// 登录校对用户名密码
func CheckUserNameAndPassword(c *gin.Context) {
	// session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	u := &entity.User{
		Username: username,
		Password: password,
	}

	accessed := userService.Login(u)
	resultDatas := gin.H{
		"username": username,
		"password": password,
		"accessed": accessed,
	}

	// with template files response
	c.HTML(http.StatusOK, "result.html", resultDatas)

}

// func QueryUser(c *gin.Context) {
// username := c.PostForm("username")
// users := userDao.QueryUsername(&username)
// resultDatas := gin.H{
// 	"users": users,
// }
// c.HTML(http.StatusOK, "result_list.html", resultDatas)
// }

// 注册
// func Register(c *gin.Context) {
// 	username := c.PostForm("username")
// 	password := c.PostForm("password")
// 	rePassword := c.PostForm("rePassword")

// 	if password != rePassword {
// 		var errFields []string
// 		errFields = append(errFields, "repassword")
// 		c.HTML(http.StatusBadRequest, "invalid.html", gin.H{
// 			"errFields": errFields,
// 		})
// 		return
// 	}

// 	u := &entity.User{
// 		Username: username,
// 		Password: password,
// 	}

//	// 参数校验demo，用于注册
//	err := dao2.GetValidator().Struct(*u)
//
//	var errFields []string
//	if err != nil {
//		msgs := err.(vldt.ValidationErrors)
//		for _, e := range msgs {
//			errFields = append(errFields, e.Field())
//		}
//		c.HTML(http.StatusBadRequest, "invalid.html", gin.H{
//			"errFields": errFields,
//		})
//		return
//	}
//	dao2.GetUserDao().Create(u)
//	accessed := dao2.GetUserDao().CheckUserNameAndPassword(u)
//	resultDatas := gin.H{
//		"username": username,
//		"password": password,
//		"accessed": accessed,
//	}
//	c.HTML(http.StatusOK, "result.html", resultDatas)
//}

// 	// 参数校验demo，用于注册
// 	err := factory.GetValidator().Struct(*u)

// 	var errFields []string
// 	if err != nil {
// 		msgs := err.(vldt.ValidationErrors)
// 		for _, e := range msgs {
// 			errFields = append(errFields, e.Field())
// 		}
// 		c.HTML(http.StatusBadRequest, "invalid.html", gin.H{
// 			"errFields": errFields,
// 		})
// 		return
// 	}
// 	factory.GetUserDao().Create(u)
// 	accessed := factory.GetUserDao().CheckUserNameAndPassword(u)
// 	resultDatas := gin.H{
// 		"username": username,
// 		"password": password,
// 		"accessed": accessed,
// 	}
// 	c.HTML(http.StatusOK, "result.html", resultDatas)
// }
