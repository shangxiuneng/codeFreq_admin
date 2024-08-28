-- 保存每道题目对应的公司
CREATE TABLE `t_question_statistics` (
                                         `id` int NOT NULL AUTO_INCREMENT,
                                         `question_id` int NOT NULL,
                                         `frequency` int DEFAULT '0',
                                         `last_reviewed_at` timestamp NULL DEFAULT NULL,
                                         `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                         `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci