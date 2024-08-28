-- 保存每个用户的做题记录
create table t_record(
    id int  -- 主键id
    user_id varchar(255)  -- 用户的user_id
    question_id int  -- 用户完成的题目id
    `master_level` int not null default 0,-- 掌握情况
)

