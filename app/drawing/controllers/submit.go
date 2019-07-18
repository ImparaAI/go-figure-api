package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"

	//"api/app/drawing/store"
)

type Drawing struct {
	Id int `json:"id"`
}

func Submit(c echo.Context) error {
	//grab input from context

	//put values into storage

	//respond with id of new drawing

	drawing := &Drawing{1}

	return c.JSON(http.StatusOK, drawing)
}