package middleware

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func IdExists(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		// ask storage if id is valid
		if !true {
			return c.String(http.StatusBadRequest, "The provided submission doesn't exist: " + id)
		}

		return next(c)
	}
}