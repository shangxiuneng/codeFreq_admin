-- 保存所有的题目
CREATE TABLE `t_question` (
                              `id` int NOT NULL AUTO_INCREMENT,
                              `question_id` int NOT NULL,  -- 题目id
                              `front_question_id` varchar(255) NOT NULL DEFAULT '',  -- 前端展示的题目id
                              `title` varchar(255) NOT NULL DEFAULT '',  -- 题目的标题
                              `difficulty` int NOT NULL DEFAULT '0',  -- 题目难度 1 简单 2 中等 3 困难
                              `slug_title` varchar(255) NOT NULL DEFAULT '',  -- leetcode
                              `expand` tinyint(1) DEFAULT '0',
                              `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                              `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              `deleted_at` timestamp NULL DEFAULT NULL,
                              PRIMARY KEY (`id`),
                              UNIQUE KEY `question_id` (`question_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
