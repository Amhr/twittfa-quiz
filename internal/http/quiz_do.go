package http

import (
	"WhoKnowsMeapp/config"
	"WhoKnowsMeapp/internal/app"
	"WhoKnowsMeapp/internal/template"
	"github.com/labstack/echo"
	"net/http"
)

func HandlerDoQuiz(hnm app.App) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		publicCode := ctx.Param("uid")
		quiz, found := hnm.GetQuizByPublic(publicCode)
		if !found {
			return ctx.Redirect(http.StatusFound, config.UrlFor("/quiz/404"))
		}
		return ctx.Render(http.StatusOK, "quiz/do", template.TemplateData(echo.Map{
			"title": "انجام کوئیز",
			"quiz":  quiz.Map(),
			"csrf":  ctx.Get("csrf"),
		}))
	}
}
