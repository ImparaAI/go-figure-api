package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"

	"api/app/drawing/store"
)

func FetchOne(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	return c.JSON(http.StatusOK, store.New().Get(id))
}
