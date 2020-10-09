package http

import (
	"WhoKnowsMeapp/internal/app"
	"WhoKnowsMeapp/internal/template"
	"github.com/labstack/echo"
	"net/http"
)

func HandlerQuizStartForm(hnm app.App) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		csrf := ctx.Get("csrf")
		return ctx.Render(http.StatusOK, "quiz/create", template.TemplateData(echo.Map{
			"csrf":  csrf,
			"title": "ساخت کویز جدید",
		}))
	}
}
