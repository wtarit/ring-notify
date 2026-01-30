package middleware

import (
	"api/configs"
	"api/models"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

// APIKeyAuthMiddleware validates API keys from the APIKey table
func APIKeyAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
		}

		// Expected format: "Bearer <api-key>"
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization header format")
		}

		apiKeyStr := parts[1]
		db := configs.DB()

		// Look up API key in APIKey table
		var apiKey models.APIKey
		result := db.First(&apiKey, "key = ? AND is_active = ?", apiKeyStr, true)

		if result.Error != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid API key")
		}

		// Check if API key has expired
		if apiKey.ExpiresAt != nil && apiKey.ExpiresAt.Before(time.Now()) {
			return echo.NewHTTPError(http.StatusUnauthorized, "API key expired")
		}

		// Update last used timestamp (async to avoid slowing down request)
		go func() {
			now := time.Now()
			db.Model(&models.APIKey{}).
				Where("id = ?", apiKey.ID).
				Update("last_used_at", now)
		}()

		// Store Supabase user ID in context
		c.Set("supabase_user_id", apiKey.UserID)
		c.Set("auth_type", "api_key")

		return next(c)
	}
}
