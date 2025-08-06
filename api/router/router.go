package router

import (
	"api/handler"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	userHandler := handler.NewUserHandler()
	notifyHandler := handler.NewNotifyHandler()

	e.POST("/user/create", userHandler.CreateUser)

	e.POST("/notify/call", notifyHandler.Call)
}
