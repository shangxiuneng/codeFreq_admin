package base

import (
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

// 自定义的错误

type BaseResp struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Prompt string      `json:"prompt"`
	Data   interface{} `json:"data"`
}

func NewSuccessResp(c *app.RequestContext, data interface{}) {
	resp := BaseResp{
		Code:   0,
		Msg:    "success",
		Prompt: "",
		Data:   data,
	}
	c.JSON(http.StatusOK, resp)
}
func NewErrResp(c *app.RequestContext, code int, msg, prompt string) {
	resp := BaseResp{
		Code:   code,
		Msg:    msg,
		Prompt: prompt,
	}
	c.JSON(http.StatusOK, resp)
}
