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
