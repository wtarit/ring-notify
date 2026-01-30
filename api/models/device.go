package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Device represents a device registered to receive notifications
type Device struct {
	gorm.Model
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primaryKey" example:"550e8400-e29b-41d4-a716-446655440000"`
	UserID       uuid.UUID `json:"userId" gorm:"type:uuid;index;not null" example:"550e8400-e29b-41d4-a716-446655440001"`
	FCMToken     string    `json:"fcmToken" gorm:"uniqueIndex;not null" example:"fcm-token-example"`
	DeviceName   string    `json:"deviceName" example:"My Phone"`
	DeviceType   string    `json:"deviceType" example:"android"`
	RegisteredAt time.Time `json:"registeredAt" example:"2023-01-01T00:00:00Z"`
	LastActive   time.Time `json:"lastActive" example:"2023-01-01T00:00:00Z"`
	IsActive     bool      `json:"isActive" gorm:"default:true" example:"true"`

	User User `json:"-" gorm:"foreignKey:UserID"`
}
