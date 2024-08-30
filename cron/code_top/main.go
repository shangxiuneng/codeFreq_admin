package main

import (
	"codeFreq_admin/dal"
	"codeFreq_admin/dal/mysql"
	"codeFreq_admin/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type CodeTopResp struct {
	Count    int    `json:"count"`
	List     []List `json:"list"`
	Finished []int  `json:"finished"`
}

type Leetcode struct {
	ID                 int    `json:"id"`
	FrontendQuestionID string `json:"frontend_question_id"`
	QuestionID         int    `json:"question_id"`
	Title              string `json:"title"`
	Content            string `json:"content"`
	Level              int    `json:"level"`
	SlugTitle          string `json:"slug_title"`
	Expand             bool   `json:"expand"`
}

type List struct {
	ID           int       `json:"id"`
	Value        int       `json:"value"`
	Time         time.Time `json:"time"`
	Status       bool      `json:"status"`
	NoteStatus   bool      `json:"note_status"`
	Rate         int       `json:"rate"`
	Leetcode     Leetcode  `json:"leetcode"`
	CommentCount int       `json:"comment_count"`
}

func main() {
	client := &http.Client{}

	// 设置请求的URL
	url := "https://codetop.cc/api/questions/?company=1&page=4&search=&ordering=-frequency"

	// 设置请求方法和请求体（如果没有请求体，则使用nil）
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Authorization", "Bearer your_token_here")
	//req.Header.Set("Custom-Header", "CustomValue")

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// 输出响应状态码和响应体
	//fmt.Println("Response Status Code:", resp.StatusCode)
	//fmt.Println("Response Body:", string(body))

	codeTop := CodeTopResp{}

	err = json.Unmarshal(body, &codeTop)
	if err != nil {
		fmt.Println(err)
		return
	}

	mysql.Init()
	questionClient := dal.NewQuestionClient(mysql.DB)
	var questionList []*model.TQuestion
	for _, question := range codeTop.List {
		questionList = append(questionList, &model.TQuestion{
			/*
				ID              int64     `json:"id" gorm:"id"`
				QuestionId      int       `json:"question_id" gorm:"question_id"`
				FrontQuestionId string    `json:"front_question_id" gorm:"front_question_id"`
				Title           string    `json:"title" gorm:"title"`
				Difficulty      int       `json:"difficulty" gorm:"difficulty"`
				SlugTitle       string    `json:"slug_title" gorm:"slug_title"`
				Expand          int8      `json:"expand" gorm:"expand"`
				Tags            string    `json:"tags" gorm:"tags"`
				Company         string    `json:"company" gorm:"company"`
				Freq            int64     `json:"freq" gorm:"freq"`
				LastTime        string    `json:"last_time" gorm:"last_time"`
				QuestionSet     string    `json:"question_set" gorm:"question_set"`
				CreatedAt       time.Time `json:"created_at" gorm:"created_at"`
				UpdatedAt       time.Time `json:"updated_at" gorm:"updated_at"`
				DeletedAt       time.Time `json:"deleted_at" gorm:"deleted_at"`
			*/
			QuestionId:      question.Leetcode.QuestionID,
			FrontQuestionId: question.Leetcode.FrontendQuestionID,
			Title:           question.Leetcode.Title,
			Difficulty:      question.Leetcode.Level,
			Freq:            int64(question.Value),
			LastTime:        question.Time.Format("2006-01-02 15:04:05"),
		})
	}
	err = questionClient.BatchInsert(questionList)
	if err != nil {
		fmt.Println(err)
		return
	}
}
