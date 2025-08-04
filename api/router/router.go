package router

import (
	"api/handler"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	userHandler := handler.NewUserHandler()

	e.POST("/user/create", userHandler.CreateUser)
}
