package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"

	"api/app/drawing/processing"
	"api/app/drawing/store"
	"api/app/drawing/types"
)

func Submit(c echo.Context) error {
	points, _ := c.Get("points").([]types.OriginalPoint)

	id := store.New(c.Request().Context()).Create(points)

	processing.AddToQueue(id)

	return c.JSON(http.StatusOK, Response{id})
}

type Response struct {
	Id int64 `json:"id"`
}
