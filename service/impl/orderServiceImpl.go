package impl

import (
	"math/rand"
	"time"
	"travel-agency/core/dto"
	"travel-agency/core/exception"
	"travel-agency/entity"
	"travel-agency/entity/association"
	"travel-agency/factory/dao"
	"travel-agency/util"

	"github.com/google/go-cmp/cmp"
	log "github.com/sirupsen/logrus"
)

var (
	orderDAO = dao.GetOrderDAO()
)

type OrderServiceImpl struct{}

func (OrderServiceImpl) SaveOrUpdateOrder(order *association.OrderDetailAssociation) (int, error) {
	log.Printf("args -> = %v", util.ObjectToStr(order))
	if order == nil {
		return 0, exception.Npe(nil)

	}
	// update
	// 查是否有此记录
	o := &entity.Order{OrderId: order.OrderId}
	emptyO := &entity.Order{OrderId: order.OrderId}
	db.Where("order_id=?", order.OrderId).First(o)
	// 若查有 即其他字段非0值
	if !cmp.Equal(o, emptyO) {
		orderDAO.SaveOrUpdate(order)
		return 1, nil
	}

	// 创建新的记录
	//before create to generate a orderid
	for true {
		rand.Seed(time.Now().UnixNano())
		genId := rand.Intn(999999990) + 6
		genNewOrder := &entity.Order{OrderId: genId}
		genEmptyOrder := &entity.Order{OrderId: genId}

		db.Find(genNewOrder)

		// flag为false表示id重复，需要重新gen
		flag := cmp.Equal(genNewOrder, genEmptyOrder)
		log.Printf("compare random id=%d with db datas: %v", genId, flag)
		//save

		if flag {
			//
			order.OrderId = genNewOrder.OrderId
			orderDAO.SaveOrUpdate(order)
			return 1, nil
		}

	}

	return 0, exception.Npe(nil)

}

// 逻辑删除 根据ID删除
func (OrderServiceImpl) DeleteOrder(order *entity.Order) error {
	log.Printf("args -> = %v", util.ObjectToStr(*order))
	if order == nil {
		return exception.Npe(nil)

	}

	db.Model(order).Update("delete_status", true)

	return nil
}
func (OrderServiceImpl) FindByCondition(order *entity.Order, page int, size int) dto.PageResult {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 5
	}
	w := &dto.PageResult{
		RequestSize: size,
		CurrentPage: page,
	}
	orderDAO.FindList(order, w)

	return *w
}

// 查order和detail全部
func (OrderServiceImpl) FindWithDetails(page int, size int, orderId *string, travelPathId *string, travelDetailId *string) dto.PageResult {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 5
	}
	if orderId == nil {
		empty := ""
		orderId = &empty
	}
	if travelPathId == nil {
		empty := ""
		travelPathId = &empty
	}
	if travelDetailId == nil {
		empty := ""
		travelDetailId = &empty
	}
	w := &dto.PageResult{
		RequestSize: size,
		CurrentPage: page,
	}
	orderDAO.FindOrderAndDetailList(w, orderId, travelPathId, travelDetailId)

	return *w
}
func (OrderServiceImpl) GetById(id int) (res entity.Order) {
	log.Printf("find byid = %d", id)

	db.Model(&res).Where("order_id = ?", id).
		First(&res)
	return res
}
func (OrderServiceImpl) GetOrderAndDetailById(id int) *association.OrderDetailAssociation {
	log.Printf("find byid = %d", id)

	return orderDAO.FindOrderAndDetailById(id)
}

// 变更订单支付状态
func (OrderServiceImpl) UpdatePayStatus(order *entity.Order) (int, error) {
	log.Printf("request = %v", util.ToStringJSON(*order))
	if order != nil {
		order.BeforeUpdate()
		db.Model(order).
			Update("pay_status", *order.PayStatus)
		db.Model(order).
			Update("gmt_modified", *order.GmtModified)
	}
	return 1, nil

}
