package db

import (
	"encoding/json"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type Result struct {
	gorm.Model
	QuizID       uint `gorm:"index"`
	Score        int
	PercentScore int
	DisplayName  string
	Answers      string
}

func (q *Result) GetAnswers() []echo.Map {
	var answers []echo.Map
	err := json.Unmarshal([]byte(q.Answers), &answers)
	if err != nil {
		answers = []echo.Map{}
	}
	return answers
}

func (r *Result) Map() echo.Map {
	return echo.Map{
		"id":           r.ID,
		"score":        r.Score,
		"percent":      r.PercentScore,
		"display_name": r.DisplayName,
		"answers":      r.GetAnswers(),
	}
}
