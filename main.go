package main

import (
	"os"

	"api/app"
	"api/database"
)

func main() {
	e := app.New()

	dbError := database.Initialize()

	if dbError == nil {
		e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
	} else {
		e.Logger.Fatal(dbError)
	}
}