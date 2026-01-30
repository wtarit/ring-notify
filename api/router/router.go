package router

import (
	"api/handler"
	"api/middleware"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	userHandler := handler.NewUserHandler()
	deviceHandler := handler.NewDeviceHandler()
	apiKeyHandler := handler.NewAPIKeyHandler()
	notifyHandler := handler.NewNotifyHandler()

	// User routes - anonymous user creation and API key regeneration
	e.POST("/users", userHandler.CreateUser)
	e.POST("/users/api-key", userHandler.RegenerateAPIKey, middleware.APIKeyAuthMiddleware)

	// Notify routes - requires API key authentication
	notifyGroup := e.Group("/notify")
	notifyGroup.Use(middleware.APIKeyAuthMiddleware)
	notifyGroup.POST("/call", notifyHandler.Call)

	// Device management routes - requires Supabase authentication
	deviceGroup := e.Group("/devices")
	deviceGroup.Use(middleware.SupabaseAuthMiddleware)
	deviceGroup.POST("", deviceHandler.Register)
	deviceGroup.GET("", deviceHandler.List)
	deviceGroup.PATCH("/:id", deviceHandler.Update)
	deviceGroup.DELETE("/:id", deviceHandler.Remove)

	// API key management routes - requires Supabase authentication
	apiKeyGroup := e.Group("/api-keys")
	apiKeyGroup.Use(middleware.SupabaseAuthMiddleware)
	apiKeyGroup.POST("", apiKeyHandler.Create)
	apiKeyGroup.GET("", apiKeyHandler.List)
	apiKeyGroup.DELETE("/:id", apiKeyHandler.Revoke)
}
