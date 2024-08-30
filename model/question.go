package model

import (
	"time"
)

// TQuestion undefined
type TQuestion struct {
	ID              int64      `json:"id" gorm:"id"`
	QuestionId      int        `json:"question_id" gorm:"question_id"`
	FrontQuestionId string     `json:"front_question_id" gorm:"front_question_id"`
	Title           string     `json:"title" gorm:"title"`
	Difficulty      int        `json:"difficulty" gorm:"difficulty"`
	SlugTitle       string     `json:"slug_title" gorm:"slug_title"`
	Expand          int        `json:"expand" gorm:"expand"`
	Tags            string     `json:"tags" gorm:"tags"`
	Company         string     `json:"company" gorm:"company"`
	Freq            int64      `json:"freq" gorm:"freq"`
	LastTime        string     `json:"last_time" gorm:"last_time"`
	QuestionSet     string     `json:"question_set" gorm:"question_set"`
	CreatedAt       time.Time  `json:"created_at" gorm:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at" gorm:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at" gorm:"deleted_at"`
}

// TableName 表名称
func (*TQuestion) TableName() string {
	return "t_question"
}

// QuestionListParam 查询题目列表参数
type QuestionListParam struct {
	Page      int    `query:"page"`
	Size      int    `query:"size"`
	Company   string `query:"company"`    // 公司
	Title     string `query:"title"`      // 标题或题目编号
	Difficult int    `query:"difficult"`  // 难度
	Status    int    `query:"status"`     // 状态 完成或未完成
	SetName   string `query:"set_name"`   // 题集
	Tags      string `query:"tags"`       // 标签
	LastTime  string `query:"last_time""` // 最近一次考察时间
	OrderBy   string `query:"order_by""`  // 排序规则
}

// GetQuestionListResp 查询题目列表返回参数
type GetQuestionListResp struct {
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	Count int         `json:"count"`
	List  []*Question `json:"list"`
}

type Question struct {
	QuestionId      int      `json:"question_id"`
	FrontQuestionId string   `json:"front_question_id"`
	Title           string   `json:"title"`
	Difficulty      string   `json:"difficulty"`
	SlugTitle       string   `json:"slug_title"`
	Expand          int      `json:"expand"`
	Freq            int64    `json:"freq"`             // 题目出现的频率
	LastReviewedAt  string   `json:"last_reviewed_at"` // 最近一次的考察时间
	Tags            []string `json:"tags"`             // 题目所属的标签
	Companies       []string `json:"companies"`        // 考察该题目的公司
}
