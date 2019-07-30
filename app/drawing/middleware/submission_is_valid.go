package middleware

import (
	"strings"
	"net/http"
	"github.com/labstack/echo/v4"

	apphttp "api/app/http"
	"api/app/drawing/types"
)

func SubmissionIsValid(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		input := &types.SubmitInput{}
		err := apphttp.BuildJson(c, input)

		if (err != nil) || (input.Points == nil) {
			return echo.NewHTTPError(http.StatusBadRequest, "The request is not properly formatted.")
		}

		if !strings.HasPrefix(input.Image, "data:image/png;") {
			return echo.NewHTTPError(http.StatusBadRequest, "The image is invalid.")
		}

		if len(input.Points) == 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "There needs to be at least 1 point.")
		}

		if input.Points[0].Time != 0 {
			return echo.NewHTTPError(http.StatusBadRequest, "The first point's time must be zero.")
		}

		if pointsAreSequential(input.Points) == false {
			return echo.NewHTTPError(http.StatusBadRequest, "Each point's time should be equal to or greater than the previous point.")
		}

		c.Set("image", input.Image)
		c.Set("points", input.Points)

		return next(c)
	}
}

func pointsAreSequential(points []types.OriginalPoint) bool {
	var lastPoint types.OriginalPoint

	for i := 0; i < len(points); i++ {
		if (i != 0) && (points[i].Time < lastPoint.Time) {
			return false
		}

		lastPoint = points[i]
	}

	return true
}