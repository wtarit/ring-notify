package handler

import (
	"api/configs"
	"api/notify"
	"api/user"

	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	configs.InitFirebase()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})
	e.POST("/notify/call", notify.Call)
	e.POST("/user/create", user.CreateUser)
	return e
}

func Handler(w http.ResponseWriter, r *http.Request) {
	e := Init()
	e.ServeHTTP(w, r)
}
