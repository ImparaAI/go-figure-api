package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func Submit(c echo.Context) error {
	//grab input from context

	//put values into storage

	//respond with id of new submission

	return c.String(http.StatusOK, "Your submission has been received, now wait for our thing to be done")
}