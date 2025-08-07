package util

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomBinder struct{}

func (cb *CustomBinder) Bind(i interface{}, c echo.Context) error {
	req := c.Request()
	if req.Body == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "request body is empty")
	}

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request: "+err.Error())
	}

	return nil
}
