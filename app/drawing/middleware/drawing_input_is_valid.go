package middleware

import (
	"net/http"
	"github.com/labstack/echo/v4"

	apphttp "api/app/http"
	"api/app/drawing/types"
)

func DrawingInputIsValid(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		input := &types.SubmitInput{}
		err := apphttp.BuildJson(c, input)

		if (err != nil) || (input.Points == nil) {
			return echo.NewHTTPError(http.StatusBadRequest, "The request is not properly formatted.")
		}

		if len(input.Points) == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "There needs to be at least 1 point.")
		}

		c.Set("input", input)

		return next(c)
	}
}