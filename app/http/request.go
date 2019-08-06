package http

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
)

func BuildJson(c echo.Context, inputMap interface{}) error {
	request := c.Request()
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		return err
	}

	request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return json.Unmarshal(body, inputMap)
}
