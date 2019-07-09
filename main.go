package main

import (
	"os"
	"api/app"
)

func main() {
	e := app.New()

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}