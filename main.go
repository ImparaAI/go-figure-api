package main

import (
	"github.com/labstack/echo/v4"

	"app/app/article"
)

func main() {
	e := echo.New()

	registerServices(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func registerServices(e *echo.Echo) {
	article.Register(e)
}