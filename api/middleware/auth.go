package middleware

import (
	"api/configs"
	"api/models"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization header format")
		}

		tokenStr := parts[1]
		tokenUUID, err := uuid.Parse(tokenStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token format")
		}
		db := configs.DB()
		var user models.User
		result := db.First(&user, "api_key = ?", tokenUUID)
		if result.Error != nil {
			log.Println(result.Error.Error())
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Token")
		}

		// Store user in context
		c.Set("user", &user)
		return next(c)
	}
}
