package app

import (
	"WhoKnowsMeapp/config"
	"WhoKnowsMeapp/internal/db"
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
)

func (a *App) SaveResult(quiz *db.Quiz, answers []string, displayName string) (*db.Result, error) {
	userAnswers := quiz.GetAnswers()
	questions := quiz.GetQuestions()
	checked := []map[string]interface{}{}
	score := 0
	for i, answer := range answers {
		didAnswer := answer == userAnswers[i]
		if didAnswer {
			score++
		}
		checked = append(checked, map[string]interface{}{
			"question":     questions[i]["TextFromOthers"],
			"valid_answer": userAnswers[i],
			"answer":       answer,
		})
	}
	percentScore := score * 100 / len(answers)
	j, _ := json.Marshal(checked)
	result := &db.Result{
		QuizID:       quiz.ID,
		Score:        score,
		PercentScore: percentScore,
		DisplayName:  displayName,
		Answers:      string(j),
	}
	r := a.Gorm.Create(result)
	if r.Error != nil {
		return nil, r.Error
	} else {
		return result, nil
	}
}

func (a *App) GetResults(q *db.Quiz, page int) []db.Result {
	cacheKey := fmt.Sprintf("results_%d_%d", q.ID, page)
	items := a.Redis.Get(context.Background(), cacheKey)
	var results []db.Result
	if items.Err() != nil {
		a.Gorm.Limit(config.APP_LIST_PER_PAGE+1).Offset(config.APP_LIST_PER_PAGE*(page-1)).Where("quiz_id = ?", q.ID).Order("id DESC").Find(&results)
		value, _ := json.Marshal(results)
		a.Redis.Set(context.Background(), cacheKey, string(value), config.APP_LIST_CACHE_TIME)
	} else {
		result, _ := items.Result()
		json.Unmarshal([]byte(result), &results)
	}
	return results

}

func (a *App) ResultsToMap(ress []db.Result) []echo.Map {
	m := []echo.Map{}
	for _, r := range ress {
		m = append(m, r.Map())
	}
	return m
}
