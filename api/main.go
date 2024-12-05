package main

import (
	"api/configs"
	"api/notify"
	"api/user"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatalln("Error loading .env")
	}

	configs.InitFirebase()
	configs.InitDatabase()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})
	e.POST("/notify/call", notify.Call)
	e.POST("/user/create", user.CreateUser)
	e.Logger.Fatal(e.Start(":1323"))
}
