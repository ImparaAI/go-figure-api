package middleware

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func DrawingIsValid(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !true {
			return c.String(http.StatusBadRequest, "there's a problem with your drawing")
		}

		return next(c)
	}
}