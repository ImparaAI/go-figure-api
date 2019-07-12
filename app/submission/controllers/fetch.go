package controllers

import (
	"strconv"
	"net/http"
	"github.com/labstack/echo/v4"

	"api/app/submission/store"
)

func FetchOne(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	submission, _ := store.New().Get(id)

	return c.JSON(http.StatusOK, submission)
}