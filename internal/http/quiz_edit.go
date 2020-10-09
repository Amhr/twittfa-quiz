package http

import (
	"WhoKnowsMeapp/config"
	"WhoKnowsMeapp/internal/app"
	"WhoKnowsMeapp/internal/template"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func HandlerEditQuiz(hnm app.App) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		private := ctx.Param("uid")
		if private == "" {
			return ctx.Redirect(http.StatusFound, config.UrlFor("/quiz/start"))
		}

		quiz, founded := hnm.GetQuizByPrivate(private)

		if !founded {
			return ctx.Redirect(http.StatusFound, config.UrlFor("/quiz/404"))
		}

		if quiz.Status == 5 {
			// quiz needs to be initialized
			return ctx.Render(http.StatusOK, "quiz/init", template.TemplateData(echo.Map{
				"title": "آماده سازی کوئیز",
				"quiz":  quiz.Map(),
				"csrf":  ctx.Get("csrf"),
			}))
		}

		pageParam := ctx.QueryParam("page")
		page, e := strconv.Atoi(pageParam)
		if e != nil || page < 1 {
			page = 1
		}

		results := hnm.ResultsToMap(hnm.GetResults(quiz, page))
		nextPage := false
		prevPage := false
		if len(results) > config.APP_LIST_PER_PAGE {
			nextPage = true
			results = results[:len(results)-1]
		}

		if page > 1 {
			prevPage = true
		}

		return ctx.Render(http.StatusOK, "quiz/panel", template.TemplateData(echo.Map{
			"title":     "مدیریت کوئیز",
			"quiz":      quiz.Map(),
			"csrf":      ctx.Get("csrf"),
			"results":   results,
			"has_next_page": nextPage,
			"next_page": page+1,
			"has_prev_page": prevPage,
			"prev_page": page-1,
		}))
	}
}
