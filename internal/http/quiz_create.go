package http

import (
	"WhoKnowsMeapp/config"
	"WhoKnowsMeapp/internal/app"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func HandlerCreateQuiz(hnm app.App) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		display_name := ctx.FormValue("display_name")
		if display_name == "" {
			display_name = config.APP_DEFAULT_DISPLAY_NAME
		}

		quiz := hnm.CreateQuiz(display_name)
		return ctx.Redirect(http.StatusFound, config.UrlFor(fmt.Sprintf("/quiz/edit/%s", quiz.PrivateCode)))
	}
}
