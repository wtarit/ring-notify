package service

import (
	"api/configs"
	"api/models"
	"time"

	"github.com/google/uuid"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) CreateUser(fcmToken string) *models.User {
	db := configs.DB()
	var existing models.User
	if err := db.First(&existing, "fcm_key = ?", fcmToken).Error; err == nil {
		return &existing
	}

	u := models.User{
		ID:            uuid.New(),
		APIKey:        uuid.NewString(),
		FCMKey:        fcmToken,
		UserCreated:   time.Now(),
		FCMKeyUpdated: time.Now(),
	}
	db.Create(&u)
	return &u
}

func (s *UserService) RegenerateAPIKey(userID uuid.UUID) (*models.User, error) {
	db := configs.DB()
	var user models.User
	if err := db.First(&user, "id = ?", userID).Error; err != nil {
		return nil, err
	}
	user.APIKey = uuid.NewString()
	if err := db.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
