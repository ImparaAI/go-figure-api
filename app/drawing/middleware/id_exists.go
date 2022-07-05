package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"api/app/drawing/store"
)

func IdExists(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "The request is not properly formatted.")
		}

		if store.New(c.Request().Context()).Exists(id) != true {
			return echo.NewHTTPError(http.StatusNotFound, "This drawing doesn't exist.")
		}

		return next(c)
	}
}
