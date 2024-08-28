package dal

import (
	"codeFreq_admin/model"
	"gorm.io/gorm"
)

type QuestionClient interface {
	QueryQuestionList(param model.QuestionListParam) ([]*model.TQuestion, int, error)
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
func (q *questionClient) QueryQuestionList(param model.QuestionListParam) ([]*model.TQuestion, int, error) {

	var list []*model.TQuestion
	session := q.db.Model(&model.TQuestion{})
	if param.Title != "" {
		session = session.Where("title like ? or question_id = ?",
			"%"+param.Title+"%", param.Title)
	}
	if param.Company != "" {
		session = session.Where("company like ?", "%"+param.Company+"%")
	}

	if param.Difficult != 0 {
		session = session.Where("difficulty = ?", param.Difficult)
	}

	err := session.
		Offset((param.Page - 1) * param.Size).
		Limit(param.Size).Find(&list).Error
	return list, 0, err
}
