package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"

	"api/app/submission/store"
)

func FetchOne(c echo.Context) error {
	//id := c.Param("id")

	//put values into storage

	//respond with id of new submission

	return c.String(http.StatusOK, "great")
}

func Test(c echo.Context) error {
	return c.String(http.StatusOK, store.New().Get(1))
}