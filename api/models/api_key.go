package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// APIKey represents an API key for authentication
type APIKey struct {
	gorm.Model
	ID         uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey" example:"550e8400-e29b-41d4-a716-446655440000"`
	UserID     uuid.UUID  `json:"userId" gorm:"type:uuid;index;not null" example:"550e8400-e29b-41d4-a716-446655440001"`
	Key        string     `json:"key" gorm:"uniqueIndex;not null" example:"550e8400-e29b-41d4-a716-446655440002"`
	Name       string     `json:"name" example:"ESP32 Doorbell"`
	CreatedAt  time.Time  `json:"createdAt" example:"2023-01-01T00:00:00Z"`
	ExpiresAt  *time.Time `json:"expiresAt,omitempty" example:"2024-01-01T00:00:00Z"`
	LastUsedAt *time.Time `json:"lastUsedAt,omitempty" example:"2023-01-01T00:00:00Z"`
	IsActive   bool       `json:"isActive" gorm:"default:true" example:"true"`

	User User `json:"-" gorm:"foreignKey:UserID"`
}
