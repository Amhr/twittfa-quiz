package http

import (
	"WhoKnowsMeapp/config"
	"WhoKnowsMeapp/internal/app"
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
)

func HandlerSaveAnswers(hnm app.App) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		public := ctx.Param("uid")
		displayName := ctx.FormValue("display_name")
		if displayName == "" {
			displayName = config.APP_DEFAULT_DISPLAY_NAME
		}
		if public == "" {
			return ctx.JSON(http.StatusOK, echo.Map{
				"ok":   false,
				"data": "کوئیز یافت نشد",
			})
		}
		quiz, founded := hnm.GetQuizByPublic(public)
		if !founded {
			return ctx.JSON(http.StatusOK, echo.Map{
				"ok":   false,
				"data": "کوئیز یافت نشد",
			})
		}

		if quiz.Status != 1 {
			return ctx.JSON(http.StatusOK, echo.Map{
				"ok":   false,
				"data": "کوئیز یافت نشد",
			})
		}

		answers := ctx.FormValue("answers")
		var items []string
		err := json.Unmarshal([]byte(answers), &items)
		if err != nil {
			return ctx.JSON(http.StatusOK, echo.Map{
				"ok":   false,
				"data": "مقادیر ارسالی اشتباه اند",
			})
		}
		if len(items) != quiz.Map()["count_questions"] {
			return ctx.JSON(http.StatusOK, echo.Map{
				"ok":   false,
				"data": "به همه موارد پاسخ بدهید",
			})
		}

		result, err := hnm.SaveResult(quiz, items, displayName)

		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, echo.Map{
			"ok":   true,
			"data": result.Map(),
		})

	}
}
