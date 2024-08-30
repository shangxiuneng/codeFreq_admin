-- 保存用户提交的 待审核的题目
CREATE TABLE t_contributed_questions (
id INT AUTO_INCREMENT PRIMARY KEY,  -- 主键ID，自增
question_id INT NOT NULL,  -- 题目ID
 user_id VARCHAR(255) NOT NULL,  -- 贡献人ID
company VARCHAR(255) NOT NULL,  -- 考察该题目的公司
contribution_text TEXT NOT NULL,  -- 用户提交的题目描述或内容
audit_status INT DEFAULT 1,  -- 审核状态 -1: 审核拒绝, 1: 审核中, 2: 审核通过
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- 创建时间
updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP  -- 更新时间
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
