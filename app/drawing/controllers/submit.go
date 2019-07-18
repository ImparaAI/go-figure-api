package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"

	//"api/app/drawing/store"
)

func Submit(c echo.Context) error {
	input := c.Get("input")

	//put values into storage

	//respond with id of new drawing

	return c.JSON(http.StatusOK, input)
}