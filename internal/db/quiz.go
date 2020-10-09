package db

import (
	"encoding/json"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	QuestionID  int `gorm:"index"`
	Question    QuizQuestion
	DisplayName string `gorm:"type:varchar(100)"`
	PublicCode  string `gorm:"index"`
	PrivateCode string `gorm:"index;size:256"`
	Status      int    `gorm:"type:int(1);default:5"`
}

type QuizQuestion struct {
	gorm.Model
	Questions string
	Answers   string
}

func (q *Quiz) GetAnswers() []string {
	var answers []string
	err := json.Unmarshal([]byte(q.Question.Answers), &answers)
	if err != nil {
		answers = []string{}
	}
	return answers
}

func (q *Quiz) GetQuestions() []map[string]interface{} {
	var questions []map[string]interface{}
	err := json.Unmarshal([]byte(q.Question.Questions), &questions)
	if err != nil {
		questions = []map[string]interface{}{}
	}
	return questions
}

func (q *Quiz) Map() echo.Map {
	questions := q.GetQuestions()
	return echo.Map{
		"private_code":    q.PrivateCode,
		"public_code":     q.PublicCode,
		"display_name":    q.DisplayName,
		"status":          q.Status,
		"questions":       questions,
		"count_questions": len(questions),
		"questions_json":  q.Question.Questions,
	}
}
