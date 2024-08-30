package model

import "time"

// TContributedQuestion 表示数据库中的 t_contributed_questions 表
type TContributedQuestion struct {
	ID               int       `gorm:"primaryKey;autoIncrement"`
	QuestionID       int       `gorm:"not null"`
	UserID           string    `gorm:"type:varchar(255);not null"`
	Company          string    `gorm:"type:varchar(255);not null"`
	ContributionText string    `gorm:"type:text;not null"`
	AuditStatus      int       `gorm:"default:1"` // -1: 审核拒绝, 1: 审核中, 2: 审核通过
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime"`
}

func (TContributedQuestion) TableName() string {
	return "t_contributed_questions"
}

type ContributedQuestionReq struct {
	QuestionID int    `json:"question_id"` // 题目id
	UserID     string `json:"user_id"`     // 用户的id
	Company    string `json:"company"`     // 考察公司
}
