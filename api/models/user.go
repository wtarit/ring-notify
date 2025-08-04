package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	gorm.Model
	ID            uuid.UUID `json:"id" example:"uuid"`
	APIKey        string    `json:"apiKey" example:"uuid"`
	FCMKey        string    `json:"fcmKey" example:"fcm-token-example"`
	UserCreated   time.Time `json:"userCreated" example:"2023-01-01T00:00:00Z"`
	FCMKeyUpdated time.Time `json:"fcmKeyUpdated" example:"2023-01-01T00:00:00Z"`
}
