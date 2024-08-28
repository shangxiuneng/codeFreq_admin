package service

import (
	"codeFreq_admin/dal"
	"codeFreq_admin/dal/mysql"
	"codeFreq_admin/model"
	"context"
	"fmt"
	"strings"
)

type QuestionService interface {
	GetQuestionList(param model.QuestionListParam) ([]*model.QuestionResp, error)
}

type questionService struct {
	ctx context.Context
}

func NewQuestionService(ctx context.Context) QuestionService {
	return &questionService{
		ctx: ctx,
	}
}

// GetQuestionList 获取题目列表
func (q *questionService) GetQuestionList(param model.QuestionListParam) ([]*model.QuestionResp, error) {
	questionList, _, err := dal.NewQuestionClient(mysql.DB).QueryQuestionList(param)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	/*
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
	*/
	var resp []*model.QuestionResp
	for _, question := range questionList {

		resp = append(resp, &model.QuestionResp{
			QuestionId:      question.QuestionId,
			FrontQuestionId: question.FrontQuestionId,
			Title:           question.Title,
			Difficulty:      model.DifficultMapping[question.Difficulty],
			Tags:            strings.Split(question.Tags, ","),
		})
	}

	return resp, nil
}
