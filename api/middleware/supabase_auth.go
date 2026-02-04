package middleware

import (
	"api/configs"
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

// SupabaseAuthMiddleware validates Supabase JWT tokens using JWK
func SupabaseAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
		}

		// Expected format: "Bearer <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization header format")
		}

		tokenStr := parts[1]

		// Fetch JWKS from Supabase
		ctx := context.Background()
		keySet, err := jwk.Fetch(ctx, configs.Supabase().JWKSURL)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch JWKS")
		}

		// Parse and verify JWT with JWK Set
		token, err := jwt.Parse(
			[]byte(tokenStr),
			jwt.WithKeySet(keySet),
			jwt.WithValidate(true),
		)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		// Extract Supabase user ID from claims
		var sub string
		if err := token.Get("sub", &sub); err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing sub claim")
		}

		parsedID, err := uuid.Parse(sub)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid sub claim")
		}

		// Store Supabase user ID in context
		c.Set("supabase_user_id", parsedID)
		c.Set("auth_type", "supabase")

		return next(c)
	}
}
