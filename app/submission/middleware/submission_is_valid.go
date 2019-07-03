package middleware

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func SubmissionIsValid(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !true {
			return c.String(http.StatusBadRequest, "there's a problem with your submission")
		}

		return next(c)
	}
}