package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"

	"api/app/drawing/store"
	"api/app/drawing/types"
	"api/app/drawing/processing"
)

func Submit(c echo.Context) error {
	image, _ := c.Get("image").(string)
	points, _ := c.Get("points").([]types.OriginalPoint)

	id := store.New().Create(points, string(image))

	processing.AddToQueue(id)

	return c.JSON(http.StatusOK, Response{id})
}

type Response struct {
	Id int `json:"id"`
}