package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents a user in the system
// Both authenticated (Supabase) and anonymous users use this model
// Anonymous users have SupabaseUserID = NULL
type User struct {
	gorm.Model
	ID             uuid.UUID `json:"id" gorm:"type:uuid;primaryKey" example:"550e8400-e29b-41d4-a716-446655440000"`
	SupabaseUserID *string   `json:"supabaseUserId,omitempty" gorm:"uniqueIndex" example:"550e8400-e29b-41d4-a716-446655440001"`
	Email          *string   `json:"email,omitempty" example:"user@example.com"`
	UserCreated    time.Time `json:"userCreated" example:"2023-01-01T00:00:00Z"`

	// Relationships - used by all users (authenticated and anonymous)
	Devices []Device `json:"devices,omitempty" gorm:"foreignKey:UserID"`
	APIKeys []APIKey `json:"apiKeys,omitempty" gorm:"foreignKey:UserID"`
}
