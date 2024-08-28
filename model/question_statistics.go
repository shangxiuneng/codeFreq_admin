package model

import "time"

type QuestionStatistics struct {
	ID             int        `gorm:"primaryKey;autoIncrement"`                                             // 主键id，自增
	QuestionID     int        `gorm:"not null"`                                                             // 题目id，非空
	Frequency      int        `gorm:"default:0"`                                                            // 题目出现的频率，默认值为 0
	LastReviewedAt *time.Time `gorm:"type:timestamp;default:NULL"`                                          // 题目最近的考察时间，可能为 NULL
	CreatedAt      time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`                             // 创建时间
	UpdatedAt      time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"` // 更新时间
}

func (QuestionStatistics) TableName() string {
	return "t_question_statistics"
}
