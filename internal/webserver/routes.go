package webserver

import (
	"WhoKnowsMeapp/internal/app"
	http2 "WhoKnowsMeapp/internal/http"
	"github.com/labstack/echo"
)

func setupRoutes(e *echo.Echo, ap *app.App) {
	e.GET("/quiz/start", http2.HandlerQuizStartForm(*ap))
	e.POST("/quiz/new", http2.HandlerCreateQuiz(*ap))
	e.GET("/quiz/edit/:uid", http2.HandlerEditQuiz(*ap))
	e.POST("/quiz/update/:uid", http2.HandlerUpdateQuiz(*ap))
	e.GET("/quiz/do/:uid", http2.HandlerDoQuiz(*ap))
	e.POST("/quiz/save/:uid", http2.HandlerSaveAnswers(*ap))
}
