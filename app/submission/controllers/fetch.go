package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"

	"api/database"
	"api/app/submission/store"
)

func FetchOne(c echo.Context) error {
	id := c.Param("id")

	//put values into storage

	//respond with id of new submission

	return c.String(http.StatusOK, "Here is your submission name: " + id)
}

func Test(c echo.Context) error {
	connection := database.GetConnection()

	connection.Val = "test"

	return c.String(http.StatusOK, store.New().Get())
}