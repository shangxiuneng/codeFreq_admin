package dal

import (
	"codeFreq_admin/model"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type QuestionClient interface {
	QueryQuestionList(param model.QuestionListParam) ([]*model.TQuestion, int64, error)
	BatchInsert(data []*model.TQuestion) error
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
func (q *questionClient) QueryQuestionList(param model.QuestionListParam) ([]*model.TQuestion, int64, error) {

	var list []*model.TQuestion
	session := q.db.Model(&model.TQuestion{})
	if param.Company != "" {
		session = session.Where("company like ?", "%"+param.Company+"%")
	}
	if param.Title != "" {
		session = session.Where("title like ? or question_id = ?",
			"%"+param.Title+"%", param.Title)
	}
	if param.Difficult != 0 {
		session = session.Where("difficulty = ?", param.Difficult)
	}

	if param.SetName != "" {
		session = session.Where("question_set like ?", "%"+param.SetName+"%")
	}

	if param.Tags != "" {
		session = session.Where("tags like ?", "%"+param.Tags+"%")
	}

	if param.OrderBy != "" {
		sortField := param.OrderBy
		sortOrder := "asc"
		if strings.HasPrefix(param.OrderBy, "-") {
			sortField = param.OrderBy[1:]
			sortOrder = "desc"
		}

		session = session.Order(fmt.Sprintf("%s %s", sortField, sortOrder))
	}

	err := session.
		Offset((param.Page - 1) * param.Size).
		Limit(param.Size).Find(&list).Error

	if err != nil {
		return nil, 0, err
	}

	var cnt int64
	err = session.Count(&cnt).Error
	return list, cnt, err
}

func (q *questionClient) BatchInsert(data []*model.TQuestion) error {
	return q.db.CreateInBatches(data, len(data)).Error
}
