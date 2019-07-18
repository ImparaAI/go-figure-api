package middleware

import (
	"net/http"
	"github.com/labstack/echo/v4"

	apphttp "api/app/http"
	"api/app/drawing/types"
)

func DrawingInputIsValid(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var input *types.SubmitInput
		err := apphttp.BuildJson(c, input)

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "The request is not properly formatted.")
		}

		return next(c)
	}
}