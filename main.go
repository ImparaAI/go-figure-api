package main

import (
	"github.com/labstack/echo/v4"

	"app/app/submission"
)

func main() {
	e := echo.New()

	registerServices(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func registerServices(e *echo.Echo) {
	submission.Register(e)
}