package impl

import (
	log "github.com/sirupsen/logrus"
	"travel-agency/core/dto"
	"travel-agency/core/exception"
	"travel-agency/entity"
	"travel-agency/entity/resultMap"
)

type TravelPathDetailDAOImpl struct {
	//Db *gorm.DB
}

//var db2 = dbfactory.GetGormDB()
func (cur *TravelPathDetailDAOImpl) SaveTravelPathDetail(obj *entity.TravelPathDetail) (int, error) {
	if obj == nil {
		return 0,exception.Npe(nil)
	}
	//fmt.Println("------------sss")
	if   obj.TravelDetailId>0 {
		db.Model(&entity.TravelPathDetail{}).Where("travel_detail_id = ?",obj.TravelDetailId).
			Updates(obj)
	}else {
		db.Create(obj)
	}
	return 1,nil
}

/*
视图表：


drop view if exists   v_travel_detail_current_person_count

create view  v_travel_detail_current_person_count  as
select p.*,  ifNull(sum(od.person_count ),0)  as current_person_count  from t_travel_path_detail p  left  join t_order o on  p.travel_detail_id = o.travel_detail_id

join t_order_detail od  on  od.order_id = o.order_id

where


	  p.delete_status  = false  and  o.delete_status = false


	  group by p.travel_detail_id


select * from v_travel_detail_current_person_count  where  travel_detail_id = 10

*/

func (*TravelPathDetailDAOImpl) FindListTravelPathDetail(query *entity.TravelPathDetail,page int ,size int ) (holder dto.PageResult,res []resultMap.TravelDetailWithOrderPersonCount) {
	if query.TravelPathId== nil || *query.TravelPathId<=0 {
		log.Warn("没有 travelPathId ")
		return holder,res
	}
	if page <=0 {
		page = 1
	}
	if size <=0 {
		size = 5
	}
	if query.TeamName==nil {
		query.TeamName = new(string)
	}

	//db.Table("v_travel_detail_current_person_count").Where(query)

	//ww := query.ActiveStatus
	db.Raw("" +
		"select * from v_travel_detail_current_person_count  p " +
		"\n  " +
		"" +
		"  where p.travel_path_id = ?  and  p.team_name like" +
		"" +
		" concat('%',?,'%')    and p.delete_status = false  " , query.TravelPathId,query.TeamName ).
		Order(" active_status desc , register_end_time  desc  ").
		//Count(&holder.TotalCount).
		Limit(size).
		Offset(size*(page-1)).
		Scan(&res)


	db.Table(" v_travel_detail_current_person_count ").
		Where("travel_path_id = ?  and delete_status   = false  and team_name like  concat('%',?,'%')",query.TravelPathId ,query.TeamName).
		Count(&holder.TotalCount)
	holder.Result = &res
	//db.Count(&holder.TotalCount)

	holder.CurrentPage = page
	holder.RequestSize = size
	//计算分页
    holder.CalcTotalPage()
	return holder,res


}
