package model

import (
	"database/sql"
	"time"
)

type Question struct {
	Id              int    `gorm:"column:id;type:int(11);AUTO_INCREMENT;primary_key" json:"id"`
	QuestionId      int    `gorm:"column:question_id;type:int(11);unique;NOT NULL" json:"question_id"`
	FrontQuestionId string `gorm:"column:front_question_id;type:varchar(255);NOT NULL" json:"front_question_id"`
	Title           string `gorm:"column:title;type:varchar(255);NOT NULL" json:"title"`
	// TODO
	Difficulty int          `gorm:"column:difficulty;type:varchar(50)" json:"difficulty"`
	SlugTitle  string       `gorm:"column:slug_title;type:varchar(255)" json:"slug_title"`
	Expand     int          `gorm:"column:expand;type:tinyint(1);default:0" json:"expand"`
	CreatedAt  time.Time    `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time    `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  sql.NullTime `gorm:"column:deleted_at;type:timestamp" json:"deleted_at"`
}

func (Question) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "t_question"
}

type QuestionListParam struct {
	Page int `query:"page"`
	Size int `query:"size"`
}

type QuestionResp struct {
	QuestionId      int       `json:"question_id"`
	FrontQuestionId string    `json:"front_question_id"`
	Title           string    `json:"title"`
	Difficulty      string    `json:"difficulty"`
	SlugTitle       string    `json:"slug_title"`
	Expand          int       `json:"expand"`
	Freq            int       `json:"freq"`             // 题目出现的频率
	LastReviewedAt  string    `json:"last_reviewed_at"` // 最近一次的考察时间
	Tags            []string  `json:"tags"`             // 题目所属的标签
	Companies       []string  `json:"companies"`        // 考察该题目的公司
	CreatedAt       time.Time `json:"created_at"`
}

const (
	Easy   = 1
	Middle = 2
	Hard   = 3
)

var DifficultMapping = map[int]string{
	Easy:   "容易",
	Middle: "中等",
	Hard:   "困难",
}
