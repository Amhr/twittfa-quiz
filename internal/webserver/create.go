package webserver

import (
	"WhoKnowsMeapp/internal/app"
	"WhoKnowsMeapp/internal/template"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewWebServer() {
	App, err := app.NewApp()

	if err != nil {
		fmt.Errorf("Cannot start app %v", err)
		return
	}

	e := echo.New()

	e.Renderer = template.NewTemplate()

	e.Static("/quiz/static", "static")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:token",
	}))

	// Route => handler

	setupRoutes(e, App)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
