package main

import (
	"api/notify"
	"api/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

func initEcho() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})
	e.POST("/notify/call", notify.Call)
	e.POST("/user/create", user.CreateUser)
	return e
}

func Handler(w http.ResponseWriter, r *http.Request) {
	e := initEcho()
	e.ServeHTTP(w, r)
}

func main() {
	// http.HandleFunc("/*", Handler)
	// log.Fatal(http.ListenAndServe(":1323", nil))
	e := initEcho()
	e.Logger.Fatal(e.Start(":1323"))
}
