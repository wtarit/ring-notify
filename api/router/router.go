package router

import (
	"api/handler"
	"api/middleware"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	userHandler := handler.NewUserHandler()
	notifyHandler := handler.NewNotifyHandler()

	e.POST("/user", userHandler.CreateUser)
	e.POST("/user/api-key", userHandler.RegenerateAPIKey, middleware.AuthMiddleware)

	notifyGroup := e.Group("/notify")
	notifyGroup.Use(middleware.AuthMiddleware)
	notifyGroup.POST("/call", notifyHandler.Call)
}
