package dal

import (
	"codeFreq_admin/model"
	"gorm.io/gorm"
)

type ContributeQuestionClient interface {
	Create(data *model.TContributedQuestion) error
}

type contributeQuestionClient struct {
	db *gorm.DB
}

func NewContributeQuestionClient(db *gorm.DB) ContributeQuestionClient {
	return &contributeQuestionClient{
		db: db,
	}
}

func (c *contributeQuestionClient) Create(data *model.TContributedQuestion) error {
	err := c.db.Model(&model.TContributedQuestion{}).Create(data).Error
	return err
}
