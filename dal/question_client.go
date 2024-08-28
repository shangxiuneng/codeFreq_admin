package dal

import (
	"codeFreq_admin/model"
	"gorm.io/gorm"
)

type QuestionClient interface {
	QueryQuestionList(page, size int) ([]*model.Question, int, error)
}

type questionClient struct {
	db *gorm.DB
}

func NewQuestionClient(db *gorm.DB) QuestionClient {
	return &questionClient{
		db: db,
	}
}

// QueryQuestionList 查询题目列表
func (q *questionClient) QueryQuestionList(page, size int) ([]*model.Question, int, error) {

	var list []*model.Question

	err := q.db.Model(&model.Question{}).
		Offset((page - 1) * size).
		Limit(size).Find(&list).Error
	return list, 0, err
}
