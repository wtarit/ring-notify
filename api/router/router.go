package router

import (
	"api/handler"
	"api/middleware"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	deviceHandler := handler.NewDeviceHandler()
	apiKeyHandler := handler.NewAPIKeyHandler()
	notifyHandler := handler.NewNotifyHandler()

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
