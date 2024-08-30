package handler

import (
	"codeFreq_admin/base"
	"codeFreq_admin/model"
	"codeFreq_admin/service"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
)

func ContributeQuestion(ctx context.Context, c *app.RequestContext) {
	param := model.ContributedQuestionReq{}

	err := c.Bind(&param)
	if err != nil {
		return
	}

	fmt.Println(param)

	err = service.NewQuestionService(ctx).ContributeQuestion(param)
	if err != nil {
		base.NewErrResp(c, 5000, "", "")
		return
	}

	base.NewSuccessResp(c, nil)
}
