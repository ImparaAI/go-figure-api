package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"

	"api/app/drawing/store"
)

func FetchRecent(c echo.Context) error {
	return c.JSON(http.StatusOK, store.New().GetRecent())
}