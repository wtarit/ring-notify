package handler

import (
	"api/notify"
	"api/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitEcho() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})
	e.POST("/notify/call", notify.Call)
	e.POST("/user/create", user.CreateUser)
	return e
}

func Handler(w http.ResponseWriter, r *http.Request) {
	e := InitEcho()
	e.ServeHTTP(w, r)
}
