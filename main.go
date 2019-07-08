package main

import (
	"os"
	"github.com/labstack/echo/v4"

	"app/app/submission"
)

func main() {
	e := echo.New()

	registerServices(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}

func registerServices(e *echo.Echo) {
	submission.Register(e)
}