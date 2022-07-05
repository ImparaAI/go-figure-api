package health

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Register(e *echo.Echo) {
	e.GET("health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
}
