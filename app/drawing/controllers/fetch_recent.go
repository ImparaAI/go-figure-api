package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"

	"api/app/drawing/store"
)

func FetchRecent(c echo.Context) error {
	return c.JSON(http.StatusOK, store.New().GetRecent())
}
