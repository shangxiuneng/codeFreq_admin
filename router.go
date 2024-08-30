package main

import (
	"codeFreq_admin/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func registerRouter(h *server.Hertz) {
	h.GET("/question_list", handler.GetQuestionList)
	h.POST("/contribute_question", handler.ContributeQuestion)
}
