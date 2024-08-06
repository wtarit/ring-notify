package notify

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Call(c echo.Context) error {
	return c.String(http.StatusOK, "Called")
}
