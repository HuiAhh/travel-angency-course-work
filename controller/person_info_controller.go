package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"travel-agency/core/dto"
	"travel-agency/entity/query"
	"travel-agency/entity/resultMap"
	_ "travel-agency/service"
	"travel-agency/util"
	_ "travel-agency/factory/service"
)

//// 用户信息添加
//
//type PersonInfoController struct {
//
//
//}


//func (l *PersonInfoController) HandleUpdatePersons(c *gin.Context,dto dto.PersonInfoDTO) {
//	if dto.OrderDetailId>0 {
//		//mock 数据
//		util.RenderJson(c,nil)
//	}
//}
//
//var (
//	personInfoService service.PersonInfoService = sf.GetPersonInfoService()
//)

//person_info/edit?orderId=11
//处理请求
func HandleRequestEdit(c *gin.Context ){
	orderId := c.Query("orderId")
	log.Printf("orderId=%s",orderId)
	util.RenderHtml(c,"personEdit.html",gin.H{
		"orderId": orderId,
	})
}


//person_info/find_by_order?orderId
func JsonResultPersonInfo(c *gin.Context) {
	if orderId,er:= util.ToInt(c.Query("orderId")); er== nil && orderId>0 {
		util.RenderJson(c, dto.Ok(personInfoService.FindByOrderId(orderId)))
		return
	}
	util.RenderJson(c,dto.Fail(nil,"数据解析错误，没有orderID"))
}

//person_info/update
//更新数据
func HandleUpdateJsonPersonInfo(c *gin.Context) {
	postData := resultMap.PersonInfoWithOrder{}
	if bindErr := util.StrToJson(c,&postData); bindErr== nil {
		validErr := validator.Struct(&postData)
		if validErr==nil {
			log.Println("personInfo 数据正确")
			//TODO: 更新用户信息接口
			err := personInfoService.SavePersonInfo(&postData)
			if err ==nil {
				util.RenderJson(c,dto.Ok(nil))
				return
			}else {
				util.RenderJson(c,dto.Fail(err,"保存错误"))
			}
			//personInfoService.DeleteAllPersonInfoByOrderId(postData.OrderId)
			//personInfoService.FindByExpression()
		}
		util.RenderJson(c,dto.Fail("接口数据异常",""))
	}

}


func FindPersonsByOrder(c *gin.Context) {
	search := query.PersonInfoSearch{}

	if err :=c.ShouldBindQuery(&search);err==nil {
		page,size := util.GetPageAndSize(c)
		log.Println("request =",&search)
		res,_ := personInfoService.FindByOrderInfo(search,page,size)
		log.Println("result= ",&res)
		util.RenderHtml(c,"personSearch.html",gin.H{
			"res":res,
			"page":page,
			"size":size,
			"search":search,
		})

	}else {
		log.Println("出现异常 ",err)
		util.RenderHtml(c,"404.html",nil)
	}
}