package main

import (
	"api/configs"
	"api/notify"
	"api/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.InitFirebase()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})
	e.POST("/notify/call", notify.Call)
	e.POST("/user/create", user.CreateUser)
	e.Logger.Fatal(e.Start(":1323"))
}
