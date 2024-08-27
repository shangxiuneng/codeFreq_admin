-- 保存所有的题目
CREATE TABLE t_question (
id INT AUTO_INCREMENT PRIMARY KEY,  -- 主键id
question_id INT NOT NULL UNIQUE,    -- leetcode中对应的题目id，并设置唯一索引
front_question_id varchar(255) not null default "",
title VARCHAR(255) NOT NULL,        -- 标题
difficulty VARCHAR(50),             -- 题目难度
slug_title VARCHAR(255),            -- 直接跳转到leetcode的网站
expand BOOLEAN DEFAULT FALSE,       -- 是否有oj或者是否为拓展题
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- 创建时间
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  -- 更新时间
deleted_at TIMESTAMP NULL  -- 删除时间，软删除
);
