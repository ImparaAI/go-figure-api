package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func FetchOne(c echo.Context) error {
	id := c.Param("id")

	//put values into storage

	//respond with id of new submission

	return c.String(http.StatusOK, "Here is your great submission: " + id)
}