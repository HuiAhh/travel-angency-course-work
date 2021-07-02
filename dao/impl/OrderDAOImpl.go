package impl

import (
	"errors"
	"time"
	"travel-agency/core/dto"
	"travel-agency/entity"
	"travel-agency/entity/association"
	"travel-agency/util"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type OrderDAOImpl struct {
}

func (*OrderDAOImpl) SaveOrUpdate(obj *association.OrderDetailAssociation) (int, error) {
	if obj == nil {

		return 0, errors.New("空指针异常")

	}
	log.Printf("args ->  xx %v", util.ObjectToStr(obj))
	// todo 2021-5-24 4:49 区分save 跟 update 不依赖orderId: 直接save
	// take apart order and detail
	order := &entity.Order{
		OrderId: obj.OrderId,
		// TravelPathId:   obj.TravelPathId,
		TravelDetailId: *obj.TravelDetailId,
		PayStatus:      obj.PayStatus,
		DeleteStatus:   obj.DeleteStatus,
		UpdateByUserid: obj.UpdateByUserid,
	}

	// 根据旅行团id TravelDetailId 查路线ID TravelPathId in t_travel_path_detail
	tempTravelGroup := &entity.TravelPathDetail{}
	db.Where("travel_detail_id=?", order.TravelDetailId).First(tempTravelGroup)
	log.Printf("1.from TravelDetailId=%v query result: TravelPathId:---->%v",
		tempTravelGroup.TravelDetailId,
		tempTravelGroup.TravelPathId)
	order.TravelPathId = *tempTravelGroup.TravelPathId
	log.Printf("2.from TravelDetailId=%v query result: TravelPathId:---->%v",
		order.TravelDetailId,
		order.TravelPathId)

	//init pay status & delete status
	var falsePtr *bool = new(bool)
	*falsePtr = false
	order.PayStatus = falsePtr
	order.DeleteStatus = falsePtr

	detail := &entity.OrderDetail{
		OrderDetailId:          obj.OrderDetailId,
		PersonCount:            obj.PersonCount,
		ChildCount:             obj.ChildCount,
		OrderId:                &obj.OrderId,
		OrderDetailDescription: obj.OrderDetailDescription,
	}
	db.Save(order)
	db.Save(detail)
	return 1, nil

}
func (*OrderDAOImpl) FindList(query *entity.Order, result *dto.PageResult) {
	if query == nil || result == nil {
		return
	}
	var res []entity.Order

	db.
		Where(query).
		Offset((result.CurrentPage - 1) * result.RequestSize).
		Limit(result.RequestSize).
		Order("gmt_create desc").
		Find(&res)
	db.Model(query).Count(&result.TotalCount)
	//总页数
	result.TotalPage = ((result.TotalCount - 1) / result.RequestSize) + 1

	result.Result = res
	//db.Model(query).Find(query).Limit()
}

// order-detail关联查 仅仅根据 1.order id 精确查询 2.path id与detail id
func (OrderDAOImpl) FindOrderAndDetailList(result *dto.PageResult, orderId *string, travelPathId *string, travelDetailId *string) {
	if result == nil {
		return
	}

	var res []orderDetail
	const commonQuerySql = "select * from t_order o " +
		"left join t_order_detail od on o.order_id=od.order_id "

	const commonCountSql = "select count(*) as v_count from t_order o " +
		"left join t_order_detail od on o.order_id=od.order_id "

	type tCount struct {
		VCount int
	}
	ret_count := &tCount{}

	var query *gorm.DB
	// 匹配where 子句
	// 8个判断
	//000
	if *orderId == "" && *travelPathId == "" && *travelDetailId == "" {
		query = db.Raw(commonQuerySql +
			"where delete_status=false")

		// count
		db.Raw(commonCountSql +
			"where delete_status=false").Scan(ret_count)

		//001
	} else if *orderId == "" && *travelPathId == "" && *travelDetailId != "" {
		query = db.Raw(commonQuerySql+
			"where delete_status=false and o.travel_detail_id=?", travelDetailId)

		// count
		db.Raw(commonCountSql+
			"where delete_status=false and o.travel_detail_id=?", travelDetailId).Scan(ret_count)

		//010
	} else if *orderId == "" && *travelPathId != "" && *travelDetailId == "" {
		query = db.Raw(commonQuerySql+
			"where delete_status=false and o.travel_path_id=?", travelPathId)

		//count
		db.Raw(commonCountSql+
			"where delete_status=false and o.travel_path_id=?", travelPathId).Scan(ret_count)

		//011
	} else if *orderId == "" && *travelPathId != "" && *travelDetailId != "" {
		query = db.Raw(commonQuerySql+
			"where delete_status=false and o.travel_detail_id=? and o.travel_path_id=?",
			travelDetailId, travelPathId)

		//count
		db.Raw(commonCountSql+
			"where delete_status=false and o.travel_detail_id=? and o.travel_path_id=?",
			travelDetailId, travelPathId).
			Scan(ret_count)
		//100
	} else if *orderId != "" && *travelPathId == "" && *travelDetailId == "" {
		query = db.Raw(commonQuerySql+
			"where delete_status=false and o.order_id=?", orderId)

		//count
		db.Raw(commonCountSql+
			"where delete_status=false and o.order_id=?", orderId).
			Scan(ret_count)
		//101 一起查出来 or
	} else if *orderId != "" && *travelPathId == "" && *travelDetailId != "" {
		query = db.Raw(commonQuerySql+
			"where delete_status=false and (o.order_id=? or o.travel_detail_id=?)",
			orderId, travelDetailId)

		//count
		db.Raw(commonCountSql+
			"where delete_status=false and (o.order_id=? or o.travel_detail_id=?)",
			orderId, travelDetailId).
			Scan(ret_count)
		//110 一起查出来 or
	} else if *orderId != "" && *travelPathId != "" && *travelDetailId == "" {
		query = db.Raw(commonQuerySql+
			"where delete_status=false and (o.order_id=? or o.travel_path_id=?)",
			orderId, travelPathId)

		//count
		db.Raw(commonCountSql+
			"where delete_status=false and (o.order_id=? or o.travel_path_id=?)",
			orderId, travelPathId).
			Scan(ret_count)

		//111 一起查出来 1 or (2 and 3)
	} else {
		query = db.Raw(commonQuerySql+
			"where delete_status=false and (o.order_id=? or (o.travel_path_id=? and o.travel_detail_id=?))",
			orderId, travelPathId, travelDetailId)

		//count
		db.Raw(commonCountSql+
			"where delete_status=false and (o.order_id=? or (o.travel_path_id=? and o.travel_detail_id=?))",
			orderId, travelPathId, travelDetailId).
			Scan(ret_count)
	}

	query.Offset((result.CurrentPage - 1) * result.RequestSize).
		Limit(result.RequestSize).
		Order("gmt_create desc").
		Scan(&res)

	// 查记录条数
	// db.Model(&entity.Order{}).Count(&result.TotalCount)
	// query.Count(&result.TotalCount)

	result.TotalCount = ret_count.VCount

	//总页数
	result.TotalPage = ((result.TotalCount - 1) / result.RequestSize) + 1

	length := len(res)

	// res[i]是地址
	for i := 0; i < length; i++ {
		// temp struct
		// gorm tag 坑：不解析私有结构体字段
		type tName struct {
			Username *string `gorm:"username"` //not work
		}
		type tPath struct {
			TravelPathName *string
		}
		type tTeam struct {
			BeginTime *time.Time
			TeamName  *string
		}
		// only a field struct vars
		tempName := tName{}
		tempPath := tPath{}
		tempTeam := tTeam{}

		// 1+n query: username
		db.Raw("select username from t_order o "+
			"left join t_user u on o.update_by_userid=u.id "+
			"where o.update_by_userid=?", res[i].UpdateByUserid).Scan(&tempName)

		// 1+n query: path name
		db.Raw("select travel_path_name from t_order o "+
			"left join t_travel_path tp on o.travel_path_id=tp.travel_path_id "+
			"where o.travel_path_id=?", res[i].TravelPathId).Scan(&tempPath)

		// 1+n query: begin time
		db.Raw("select begin_time, team_name from t_order o "+
			"left join t_travel_path_detail tpd on o.travel_detail_id=tpd.travel_detail_id "+
			"where o.travel_detail_id=?", res[i].TravelDetailId).Scan(&tempTeam)

		res[i].AccessUserName = tempName.Username
		res[i].TravelPathName = tempPath.TravelPathName
		res[i].BeginTime = tempTeam.BeginTime
		res[i].TeamName = tempTeam.TeamName

	}

	result.Result = res
}

// 根据id查order-detail
func (*OrderDAOImpl) FindOrderAndDetailById(id int) *association.OrderDetailAssociation {
	res := &association.OrderDetailAssociation{}
	db.Raw("select * from t_order o "+
		"left join t_order_detail od on o.order_id=od.order_id "+
		"where delete_status=false and o.order_id=?", id).
		Scan(res)
	return res
}

// 关联查询用
type orderDetail struct {
	//detail
	OrderDetailId          int
	PersonCount            *int
	TravelPathId           int
	OrderId                int
	OrderDetailDescription *string

	//order
	TravelDetailId *int
	TeamName       *string
	PayStatus      *bool
	DeleteStatus   *bool
	GmtCreate      *time.Time
	GmtModified    *time.Time
	UpdateByUserid *int

	//update by user name
	AccessUserName *string

	//travel path name
	TravelPathName *string

	//开团日期
	BeginTime *time.Time
}
