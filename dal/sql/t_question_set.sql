-- 题目合集 考察过该题目的公司 该题目的标签 均可以使用题目集合来描述
CREATE TABLE `t_question_set` (
                                  `id` int NOT NULL AUTO_INCREMENT,
                                  `set_name` varchar(255) NOT NULL,
                                  `set_enname` varchar(255) NOT NULL,
                                  `question_id` int NOT NULL,
                                  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                  `deleted_at` timestamp NULL DEFAULT NULL,
                                  PRIMARY KEY (`id`),
                                  KEY `question_id` (`question_id`),
                                  CONSTRAINT `t_question_set_ibfk_1` FOREIGN KEY (`question_id`) REFERENCES `t_question` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci