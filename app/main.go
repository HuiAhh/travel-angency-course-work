package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
	"travel-agency/core/authc"
	"travel-agency/util"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"

	"travel-agency/controller"

	log "github.com/sirupsen/logrus"
	//
	//swaggerFiles "github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"
	//
	//
	//_ "travel-agency/app/docs"
)

//参考gin  文档：https://gin-gonic.com/zh-cn/docs/examples/html-rendering/
//Lyr
func setPath(r *gin.Engine) {
	//0.首页员工操作界面
	//r.Any("/index", controller.Index)
	r.Any("/index", controller.Index)
	r.Any("/", controller.Index)

	//1.旅游路线模块 --> 表格
	r.Any("/travel_table", controller.TravelTable)
	//添加旅游后路线
	r.Any("/travel_edit", controller.FormEdit)
	//更新旅游路线
	r.POST("/travel_path/save", controller.SaveTravelPathJson)
	//更新旅游路线 上下架状态
	r.POST("/update_travel_path", controller.UpdateTravelPath)
	//路线逻辑删除
	r.POST("/travel_path/delete", controller.DeleteByTravelPathId)
	//删除旅游路线
	r.POST("/del_travel_path", controller.DeleteTravelPath)

	r.POST("file/upload", controller.UploadFile)
	r.GET("/travelPath/preview", controller.PreviewTravelPathId)

	//添加旅游团
	r.POST("/travel_detail/add", controller.AddTravelPathDetail)
	r.GET("/travel_detail/query", controller.QueryTravelDetail)

	//<<<<<<< HEAD
	//	r.GET("/travel_detail/edit",controller.EditTravelDetail)
	//	r.GET("/travel_detail/detail",controller.PreviewTravelDetail)

	//申请人添加
	r.GET("/person_info/edit", controller.HandleRequestEdit)
	//获取申请人信息， 根据订单ID ？orderId=?
	r.GET("/person_info/find_by_order", controller.JsonResultPersonInfo)
	//更新申请人信息，通过订单ID ，更新申请人信息
	r.POST("/person_info/update", controller.HandleUpdateJsonPersonInfo)

	r.GET("/travel_detail/edit", controller.EditTravelDetail)
	r.GET("/travel_detail/detail", controller.PreviewTravelDetail)
	r.GET("/person_info/tables", controller.FindPersonsByOrder)

	//TODO: 这里是分界
	//------------------------------------------

	//用户订单模块

	//=======

	//2.用户订单模块 -->表格页
	r.Any("/order_table", controller.OrderTable)
	//订单创建——编辑页
	r.Any("/order_edit", controller.OrderEditPage)
	// 订单添加、更新
	r.POST("/order/save", controller.SaveOrderJson)
	//更新订单 支付状态
	r.POST("/update_order", controller.UpdateOrderPayStatus)
	//订单逻辑删除
	r.POST("/order/delete", controller.DeleteByOrderId)

	//订单支付状态更改为'已支付'
	r.POST("/order/paid", controller.UpdateOrderPayStatus)
	// 撤单：也就是逻辑删除 跟/order/delete 是同一个
	r.POST("/order/refund", controller.DeleteByOrderId)

	//3.费用管理
	//手续费策略管理
	r.Any("/pay_fee_table", controller.PayFeeTable)
	//未支付旅游团订单
	r.Any("/non_paid_table", controller.NonPaidTable)
	//已支付旅游团订单
	r.Any("/paid_table", controller.PaidTable)

	// 交款单 order id
	r.Any("/pay/doc", controller.PayDocument)
	// 确认书 order id
	r.Any("/pay/confirmation", controller.PayConfirmation)

	//4.报表
	//每日收支
	r.Any("/daily_income", controller.DailyIncome)
	//每日开团
	r.Any("/daily_group", controller.DailyGroup)

	//5.旅行团游客管理

}

//func swaggerConfig(r *gin.Engine)  {
//	url := ginSwagger.URL("http://localhost:9888/swagger/doc.json")
//	r.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler ,url ) )
//}

//HUI
func setPath2(r *gin.Engine) {

	// // controller
	// r.POST("/login", controller.CheckUserNameAndPassword)

	// r.POST("/query", controller.QueryUser)
	// r.POST("/register", controller.Register)
}

func configLoggerLevel() {
	log.SetLevel(log.InfoLevel)
}

func setFunc(r *gin.Engine) {
	r.SetFuncMap(template.FuncMap{
		"formatTime": func(a *time.Time) string {
			if a == nil {
				return ""
			}
			return a.Format("2006-01-02 15:04:05")
		},

		"ratioPercent": func(f *float64) string {
			if f == nil {
				return ""
			}
			percentDigit := (*f) * 100
			return fmt.Sprintf("%.2f%%", percentDigit)
		},
		"moneyTwoDigitsPoint": func(f *float64) string {
			if f == nil {
				return ""
			}
			return fmt.Sprintf("￥%.2f", *f)
		},
		"no": func(a int) int {
			return a + 1
		},
		"labelDate": func(a time.Time) string {
			return a.Format("2006年01月02日")
		},
		"nonZero": func(w int) string {
			if w == 0 {
				return ""
			}
			return strconv.Itoa(w)
		},
	})
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /
func main() {
	//配置 logger 隔离级别
	configLoggerLevel()
	// init
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	//注册登录拦截
	LoginRegister(r)
	//设置登录路径
	setPath(r)
	//swaggerConfig(r)

	//自定义函数
	setFunc(r)
	//r.Use(sessions.Sessions("goSessionID", store))
	//git
	// url pattern

	// r.LoadHTMLGlob("../templates/**/*")
	r.LoadHTMLGlob("./templates/**/*")

	r.Static("/assets", "./assets")
	//设置 public 资源文件

	//r.StaticFS("/public/css",http.Dir("/css"))
	//r.StaticFS("/public/js",http.Dir("/js"))
	r.StaticFS("/static", http.Dir("./static"))
	r.StaticFS("/public", http.Dir("./public/Light-Year-Admin-Using-Iframe-v4"))
	r.StaticFile("/index.html", "/static/index.html")
	//r.StaticFile("/", "../static/index.html")
	r.StaticFile("/login.html", "/static/login.html")

	r.Run(":9889")
}

//注册登录组件
func LoginRegister(r *gin.Engine) {
	store := cookie.NewStore([]byte("travel_system...."))
	//sessionID
	r.Use(sessions.Sessions("ivsessionID", store))
	filter := &authc.LoginFilter{}
	//注册拦截器
	r.Use(filter.Authorize())

	r.GET("/free/code", authc.GenerateCode)
	r.GET("/login", func(c *gin.Context) {
		util.RenderHtml(c, "login.html", nil)
	})
	r.POST("/login", authc.HandleLogin)
	r.GET("/logout", authc.Logout)

	r.GET("/changePwd", func(context *gin.Context) {
		util.RenderHtml(context, "changePwd.html", nil)
	})
	r.POST("/changePwd", func(c *gin.Context) {
		authc.ChangePassword(c)
	})
}
