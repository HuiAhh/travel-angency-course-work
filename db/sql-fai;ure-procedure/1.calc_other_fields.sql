DROP PROCEDURE if EXISTS calc_other_fields;
DELIMITER //
-- 常用查询字段
-- 判断计算的是哪种v_days_diff : diff_type=0 算定金 diff_type=1 算退款
CREATE PROCEDURE calc_other_fields(

    inout v_person_per int,
    inout v_child_per int,
    inout v_person_count int,
    inout v_child_count int,
    inout v_begin_time datetime,
    inout v_days_diff int,
    inout v_cancel_pay_ratio decimal(4,2),
    inout v_pay_ratio decimal(4,2),
    inout v_diff_pattern int,
    in diff_type int,
    in v_travel_path_id int,
    in v_order_id int,
    in v_travel_detail_id int,
    in in_day_date datetime
)
BEGIN 
    --单记录返回
    --1.查客单价 1.2.3.
    select per_should_pay, child_should_pay
    into v_person_per, v_child_per
    from t_travel_path 
    where travel_path_id=v_travel_path_id;
    --2.查人数 1.2.3. 连表查
    select person_count, child_count
    into v_person_count, v_child_count
    from t_order o left join t_order_detail od
    on o.order_id=od.order_id
    where o.order_id=v_order_id;
    --3.查出发日 1.2.3.
    select begin_time
    into v_begin_time
    from t_travel_path_detail
    where travel_detail_id=v_travel_detail_id;

    -- 判断计算的是哪种v_days_diff : diff_type=0 算定金、余款[支付] diff_type=1 算退款
    if diff_type=1 then
        --4.当天距离出发日、算退款：cancel_pay_ratio 2.
        select DATEDIFF(v_begin_time, in_day_date) into v_days_diff;
    else
        --5.出发日距离订单创建时间、算定金、余款：pay_ratio 1.3.
        select DATEDIFF(v_begin_time,v_gmt_create) into v_days_diff;
    end if;

    --6.算天数差范围、匹配区间 1.2.3.
    if v_days_diff < 30 then
        set v_diff_pattern = 10;
    end if;
    if v_days_diff < 60 then
        set v_diff_pattern = 30;
    else
        set v_diff_pattern = 60;
    end if;
			
    --7.退款手续费、定金费率 合并2个语句 1.2.3.
    select cancel_pay_ratio, pay_ratio
    into v_cancel_pay_ratio, v_pay_ratio
    from t_money_should_pay_fee
    where distance_time_days=v_diff_pattern;
end;
//
DELIMITER ;
