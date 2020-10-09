package http

import (
	"WhoKnowsMeapp/internal/app"
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
)

func HandlerUpdateQuiz(hnm app.App) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		private := ctx.Param("uid")
		if private == "" {
			return ctx.JSON(http.StatusOK, echo.Map{
				"ok":   false,
				"data": "کوئیز یافت نشد",
			})
		}

		quiz, founded := hnm.GetQuizByPrivate(private)
		if !founded {
			return ctx.JSON(http.StatusOK, echo.Map{
				"ok":   false,
				"data": "کوئیز یافت نشد",
			})
		}

		if quiz.Status != 5 {
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

		quiz.Question.Answers = answers
		quiz.Status = 1
		hnm.UpdateQuiz(quiz)

		return ctx.JSON(http.StatusOK,echo.Map{
			"ok":true,
			"data":"کوئیز با موفقیت آپدیت شد",
		})

	}
}
