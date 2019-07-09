package app

import (
	"github.com/labstack/echo/v4"
	"api/app/submission"
)

func New() (*echo.Echo) {
	e := echo.New()

	registerServices(e)

	return e
}

func registerServices(e *echo.Echo) {
	submission.Register(e)
}