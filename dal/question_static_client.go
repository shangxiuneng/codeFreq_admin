package dal

import (
	"codeFreq_admin/model"
	"gorm.io/gorm"
)

type QuestionStaticClient interface {
	QuestionStaticListByID(questionIDs []int) ([]*model.QuestionStatistics, error)
}

type questionStaticClient struct {
	db *gorm.DB
}

func NewQuestionStaticClient(db *gorm.DB) QuestionStaticClient {
	return &questionClient{
		db: db,
	}
}

func (q *questionClient) QuestionStaticListByID(questionIDs []int) ([]*model.QuestionStatistics, error) {
	var list []*model.QuestionStatistics

	err := q.db.Model(&model.QuestionStatistics{}).
		Where("question_id in (?)", questionIDs).Find(&list).Error

	return list, err
}
