CREATE TABLE apex_session (
	session_id varchar(10) comment '赛季id',
	session_name varchar(100) comment '赛季名称',
	session_start_time timestamp comment '赛季开始时间',
	session_end_time timestamp comment '赛季结束时间',
	session_desc varchar(500) comment '赛季说明',
	primary key(session_id)
) comment 'apex赛季';