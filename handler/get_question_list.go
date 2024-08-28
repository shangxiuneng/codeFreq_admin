package handler

import (
	"codeFreq_admin/base"
	"codeFreq_admin/model"
	"codeFreq_admin/service"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func GetQuestionList(ctx context.Context, c *app.RequestContext) {

	param := model.QuestionListParam{}

	err := c.Bind(&param)

	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"message": "参数错误",
		})
		return
	}

	if param.Page < 0 {
		param.Page = 1
	}

	if param.Size < 0 || param.Size > 50 {
		param.Size = 50
	}

	fmt.Println(param)
	data, err := service.NewQuestionService(ctx).GetQuestionList(param)
	if err != nil {
		base.NewErrResp(c, 5000, err.Error(), "内部服务错误,联系开发人员解决")
		return
	}
	base.NewSuccessResp(c, data)
	return
}
