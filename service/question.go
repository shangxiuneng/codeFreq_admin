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
	GetQuestionList(param model.QuestionListParam) (*model.GetQuestionListResp, error)
	ContributeQuestion(param model.ContributedQuestionReq) error
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
func (q *questionService) GetQuestionList(param model.QuestionListParam) (*model.GetQuestionListResp, error) {
	questionList, count, err := dal.NewQuestionClient(mysql.DB).QueryQuestionList(param)
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
	var resp []*model.Question
	for _, question := range questionList {

		resp = append(resp, &model.Question{
			QuestionId:      question.QuestionId,
			FrontQuestionId: question.FrontQuestionId,
			Title:           question.Title,
			Difficulty:      model.DifficultMapping[question.Difficulty],
			SlugTitle:       question.SlugTitle,
			Expand:          question.Expand,
			Freq:            question.Freq,
			LastReviewedAt:  question.LastTime,
			Tags:            strings.Split(question.Tags, ","),
			Companies:       strings.Split(question.Company, ","),
		})
	}

	getQuestionListResp := &model.GetQuestionListResp{
		Page:  param.Page,
		Size:  param.Size,
		Count: int(count),
		List:  resp,
	}

	return getQuestionListResp, nil
}

func (q *questionService) ContributeQuestion(param model.ContributedQuestionReq) error {
	// 查看用户在规定时间内是否重复贡献题目
	// 用户是否对审核中的题目进行重复贡献
	// 将题目信息写入到审核表中
	return nil
}

func (q *questionService) RevokeQuestionApproval() error { return nil }

func (q *questionService) ApproveQuestion() error {
	/*
			修改审核表中的状态
			将题目信息写入题目表
		TODO 题目审核通过后 应该要通知用户
	*/
	return nil
}

// GetAuditQuestion 获取审核题目列表
func (q *questionService) GetAuditQuestion() error {
	return nil
}

// ChangeQuestionStatus 修改题目的状态
func (q *questionService) ChangeQuestionStatus() error {
	return nil
}
