package ctxutil

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GetSupabaseUserID extracts the Supabase user ID from the Echo context
func GetSupabaseUserID(c echo.Context) uuid.UUID {
	userID, _ := c.Get("supabase_user_id").(uuid.UUID)
	return userID
}
