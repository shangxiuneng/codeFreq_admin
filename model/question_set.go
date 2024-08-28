package model

import "time"

type QuestionSet struct {
	ID         int        `gorm:"primaryKey;autoIncrement"`                                             // 主键id，自增
	SetName    string     `gorm:"type:varchar(255);not null"`                                           // 集合名称
	SetEnName  string     `gorm:"type:varchar(255);not null"`                                           // 集合英文名称
	QuestionID int        `gorm:"not null;index"`                                                       // 题目id，外键
	CreatedAt  time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`                             // 创建时间
	UpdatedAt  time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"` // 更新时间
	DeletedAt  *time.Time `gorm:"type:timestamp;default:NULL"`                                          // 删除时间，软删除
}

func (QuestionSet) TableName() string {
	return "t_question_set"
}
