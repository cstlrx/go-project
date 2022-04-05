

CREATE DATABASE dc;

use dc;
-- 元数据表创建
CREATE TABLE meta_info (
    id int comment'自增ID',
    gmt_create datetime comment '',
    gmt_modified datetime comment '',
    meta_type varchar(16) comment '类型：指标或维度',
    code varchar(128) comment 'code',
    name varchar(128) comment '名称',
    comment text comment '描述'
);
