package requester

import (
	"strings"
	"net/http"
	"net/http/httptest"
	"github.com/labstack/echo/v4"

	"api/app"
	"api/database"
)

func Request(method string, uri string) (*Response) {
	e := app.New()

	database.SetTestingEnvironment()
	database.Initialize()

	request := httptest.NewRequest(method, uri, nil)
	responseRecorder := httptest.NewRecorder()

	e.ServeHTTP(responseRecorder, request)

	response := &Response{responseRecorder}

	return response
}

func Get(uri string) (*Response) {
	return Request(echo.GET, uri)
}

func Post(method string, uri string) (*Response) {
	return Request(echo.POST, uri)
}


type Response struct {
	ResponseRecorder *httptest.ResponseRecorder
}

func (response *Response) Ok() (bool) {
	return response.ResponseRecorder.Code == http.StatusOK
}

func (response *Response) IsNotFound() (bool) {
	return response.ResponseRecorder.Code == http.StatusNotFound
}

func (response *Response) IsBadRequest() (bool) {
	return response.ResponseRecorder.Code == http.StatusBadRequest
}

func (response *Response) Body() (string) {
	return strings.TrimRight(response.ResponseRecorder.Body.String(), "\n")
}