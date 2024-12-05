package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID            uuid.UUID
	APIKey        string
	FCMKey        string
	UserCreated   time.Time
	FCMKeyUpdated time.Time
}
