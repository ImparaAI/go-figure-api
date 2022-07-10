package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"api/app/drawing/store"
)

func FetchOne(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	return c.JSON(http.StatusOK, store.New(c.Request().Context()).Get(id))
}
