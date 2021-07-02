DROP PROCEDURE if EXISTS calc_day_income;

--
--in->日期，按半月显示折线图
-- {{ date xxxx.xxxx }}: ￥cash

-- 存储过程
-- ret incomeByDay

-- var person child gmtcreated cancel_pay_ratio
-- where paystatus gmtmodified delete
-- link orderid orderdetailid travelpathid



-- calc ratio

-- param day
DELIMITER //
-- 计算日收入
CREATE PROCEDURE calc_day_income (in_day_date DATETIME)
BEGIN 
    --变量区
    --1.客单价
    declare v_person_per int default -1;
    declare v_child_per int default -1;
    --2.人数
    declare v_person_count int default -1;
    declare v_child_count int default -1;
    --3.中间变量
    declare v_gmt_create datetime default NOW();
    declare v_gmt_modified datetime default NOW();
    declare v_cancel_pay_ratio decimal(4,2) default 0;
    declare v_pay_ratio decimal(4,2) default 0;
    declare v_begin_time datetime default NOW();

    --链接区
    declare v_order_id int default 0;
    declare v_travel_path_id int default 0;
    declare v_order_detail_id int default 0;
    declare v_travel_detail_id int default 0;

    --计算区
    declare v_person_total decimal(10,4) default 0;
    declare v_child_total decimal(10,4) default 0;
    declare v_total_income decimal(10,4) default 0;
    declare v_days_diff int default 0;
    declare v_diff_pattern int default 0;

    declare done boolean default false;
    
    --多条记录返回

    --游标定义
    --1.查当天支付记录、+余款[]: gmt_modified=[day_date], pay_status=true, delete_status=false 
    declare cur_paid cursor for 
    select order_id, travel_path_id, travel_detail_id, gmt_create, gmt_modified 
    from t_order 
    where DATEDIFF(gmt_modified, in_day_date)=0 and pay_status=true and delete_status=false;

    --2.查当天撤单记录、-退款[]: gmt_modified=[day_date], pay_status=true, delete_status=true
    declare cur_refund cursor for 
    select order_id, travel_path_id, travel_detail_id, gmt_create, gmt_modified 
    from t_order 
    where DATEDIFF(gmt_modified, in_day_date)=0 and pay_status=true and delete_status=true;

    --3.查当天申请记录，+定金[]: gmt_create=[day_date], pay_status=false, delete_status=false
    declare cur_apply cursor for 
    select order_id, travel_path_id, travel_detail_id, gmt_create, gmt_modified 
    from t_order 
    where DATEDIFF(gmt_create, in_day_date)=0 and pay_status=false and delete_status=false;

    --4.游标中的内容执行完后将done设置为true
    DECLARE CONTINUE HANDLER FOR NOT FOUND SET done = true;

    --游标使用
    open cur_paid;
    loop_paid:loop
        fetch cur_paid into v_order_id, v_travel_path_id, v_travel_detail_id, v_gmt_create, v_gmt_modified;
        if done then leave loop_paid;end if;
        call calc_other_fields(
            v_person_per,
            v_child_per,
            v_person_count,
            v_child_count,
            v_begin_time,
            v_days_diff,
            v_cancel_pay_ratio,
            v_pay_ratio,
            v_diff_pattern,
            0,
            v_travel_path_id,
            v_order_id,
            v_travel_detail_id,
            in_day_date
        );

        --1.查当天支付记录、+余款
        set v_person_total=v_person_per*v_person_count;
        set v_child_total=v_child_per*v_child_count;
        -- 算当天加了余款部分
        set v_total_income=v_total_income+(v_person_total+v_child_total)*(1-v_pay_ratio);

    end loop loop_paid;
    -- 游标会自动在begin...end语句后关闭
    -- close cur_paid;

    open cur_refund;
    loop_refund:loop
        fetch cur_paid into v_order_id, v_travel_path_id, v_travel_detail_id, v_gmt_create, v_gmt_modified;
        if done then leave loop_refund;end if;
        call calc_other_fields(
            v_person_per,
            v_child_per,
            v_person_count,
            v_child_count,
            v_begin_time,
            v_days_diff,
            v_cancel_pay_ratio,
            v_pay_ratio,
            v_diff_pattern,
            1,
            v_travel_path_id,
            v_order_id,
            v_travel_detail_id,
            in_day_date
        );

        --2.查当天撤单记录、-退款
        set v_person_total=v_person_per*v_person_count;
        set v_child_total=v_child_per*v_child_count;
        -- 当天撤单算当天减去了 除去手续费部分
        set v_total_income=v_total_income-(v_person_total+v_child_total)*(1-cancel_pay_ratio);

    end loop loop_refund;
    -- close cur_refund;

    open cur_apply;
    loop_apply:loop
        fetch cur_paid into v_order_id, v_travel_path_id, v_travel_detail_id, v_gmt_create, v_gmt_modified;
        if done then leave loop_apply;end if;
        call calc_other_fields(
            v_person_per,
            v_child_per,
            v_person_count,
            v_child_count,
            v_begin_time,
            v_days_diff,
            v_cancel_pay_ratio,
            v_pay_ratio,
            v_diff_pattern,
            0,
            v_travel_path_id,
            v_order_id,
            v_travel_detail_id,
            in_day_date
        );

        --3.查当天申请记录，+定金
        set v_person_total=v_person_per*v_person_count;
        set v_child_total=v_child_per*v_child_count;
        -- 当天撤单算当天加定金部分
        set v_total_income=v_total_income+(v_person_total+v_child_total)*v_pay_ratio;

    end loop loop_apply;
    -- close cur_apply;



    --1.查当天支付记录、+余款[]: gmt_modified=[day_date], pay_status=true, delete_status=false 
    -- select order_id, travel_path_id, travel_detail_id, gmt_create, gmt_modified 
    -- into v_order_id, v_travel_path_id, v_travel_detail_id, v_gmt_create, v_gmt_modified
    -- from t_order 
    -- where gmt_modified=in_day_date and pay_status=true and delete_status=false;

    --2.查当天撤单记录、-退款[]: gmt_modified=[day_date], pay_status=true, delete_status=true
    -- select order_id, travel_path_id, travel_detail_id, gmt_create, gmt_modified 
    -- into v_order_id, v_travel_path_id, v_travel_detail_id, v_gmt_create, v_gmt_modified
    -- from t_order 
    -- where gmt_modified=in_day_date and pay_status=true and delete_status=true;

    --3.查当天申请记录，+定金[]: gmt_modified=[day_date], pay_status=true, delete_status=true
    -- select order_id, travel_path_id, travel_detail_id, gmt_create, gmt_modified 
    -- into v_order_id, v_travel_path_id, v_travel_detail_id, v_gmt_create, v_gmt_modified
    -- from t_order 
    -- where gmt_create=in_day_date and pay_status=false and delete_status=false;

    --总返回值
    --8.算总收
    select v_total_income;

end;
//
DELIMITER ;