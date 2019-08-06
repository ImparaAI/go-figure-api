package requester

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"strings"

	"api/app"
)

func Request(method string, uri string, payload string) *Response {
	e := app.New()

	payloadReader := strings.NewReader(payload)
	request := httptest.NewRequest(method, uri, payloadReader)
	responseRecorder := httptest.NewRecorder()

	e.ServeHTTP(responseRecorder, request)

	response := &Response{responseRecorder}

	return response
}

func Get(uri string) *Response {
	return Request(echo.GET, uri, "")
}

func Post(uri string, payload string) *Response {
	return Request(echo.POST, uri, payload)
}

type Response struct {
	ResponseRecorder *httptest.ResponseRecorder
}

func (response *Response) Ok() bool {
	return response.ResponseRecorder.Code == http.StatusOK
}

func (response *Response) IsNotFound() bool {
	return response.ResponseRecorder.Code == http.StatusNotFound
}

func (response *Response) IsBadRequest() bool {
	return response.ResponseRecorder.Code == http.StatusBadRequest
}

func (response *Response) Body() string {
	return strings.TrimRight(response.ResponseRecorder.Body.String(), "\n")
}
