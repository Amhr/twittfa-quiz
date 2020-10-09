package form_validator

import "github.com/labstack/echo"

func FormIsset(ctx echo.Context, items []string) bool {
	for _, item := range items {
		if ctx.FormValue(item) == "" {
			return false
		}
	}
	return true
}
