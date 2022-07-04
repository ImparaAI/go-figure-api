package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"api/app/drawing"
	"api/app/health"
)

func New() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	registerServices(e)

	return e
}

func registerServices(e *echo.Echo) {
	drawing.Register(e)
	health.Register(e)
}
