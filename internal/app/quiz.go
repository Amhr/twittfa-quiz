package app

import (
	"WhoKnowsMeapp/internal/db"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

func (a *App) CreateQuiz(display_name string) *db.Quiz {

	jsonQuestions, _ := json.Marshal(db.QUESTIONS)
	quiz := &db.Quiz{
		DisplayName: display_name,
		PublicCode:  uuid.New().String(),
		PrivateCode: uuid.New().String(),
		Question:    db.QuizQuestion{Questions: string(jsonQuestions), Answers: "[]"},
	}
	a.Gorm.Create(quiz)
	return quiz
}

func (a *App) GetQuizByPrivate(private string) (*db.Quiz, bool) {
	quiz, err := a.GetQuizCacheByPrivate(private)
	if err != nil {
		quiz = &db.Quiz{}
		r := a.Gorm.Joins("Question").Where("private_code=?", private).First(quiz)
		if r.Error != nil {
			return nil, false
		}
		a.UpdateQuizCache(quiz)
	}
	return quiz, true
}

func (a *App) GetQuizByPublic(public string) (*db.Quiz, bool) {
	quiz, err := a.GetQuizCacheByPublic(public)
	if err != nil {
		quiz = &db.Quiz{}
		r := a.Gorm.Joins("Question").Where("public_code=?", public).First(quiz)
		if r.Error != nil {
			return nil, false
		}
		a.UpdateQuizCache(quiz)
	}
	return quiz, true
}

func (a *App) UpdateQuizCache(quiz *db.Quiz) error {
	b, err := json.Marshal(quiz)
	if err != nil {
		return err
	}
	a.BigCache.Set(fmt.Sprintf("quiz_by_private_%s", quiz.PrivateCode), b)
	a.BigCache.Set(fmt.Sprintf("quiz_by_public_%s", quiz.PublicCode), b)
	return nil
}

func (a *App) GetQuizCacheByPrivate(private string) (*db.Quiz, error) {
	return a.GetQuizCacheBy(private, "private")
}

func (a *App) GetQuizCacheByPublic(public string) (*db.Quiz, error) {
	return a.GetQuizCacheBy(public, "public")
}

func (a *App) GetQuizCacheBy(id string, t string) (*db.Quiz, error) {
	quiz := &db.Quiz{}
	b, err := a.BigCache.Get(fmt.Sprintf("quiz_by_%s_%s", t, id))
	if err == nil {
		err = json.Unmarshal(b, quiz)
		if err != nil {
			return nil, err
		}
		return quiz, nil
	}
	return nil, err
}

func (a *App) UpdateQuiz(quiz *db.Quiz) {
	a.UpdateQuizCache(quiz)
	a.Gorm.Save(quiz)
	a.Gorm.Save(quiz.Question)
}
