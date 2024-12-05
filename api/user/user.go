package user

import (
	"api/configs"
	"api/models"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CreateUserRequest struct {
	FcmToken string `query:"fcmToken"`
}

func CreateUser(c echo.Context) error {
	var reqBody CreateUserRequest
	err := c.Bind(&reqBody)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	db := configs.DB()
	u := models.User{
		ID:            uuid.New(),
		APIKey:        uuid.NewString(),
		FCMKey:        reqBody.FcmToken,
		UserCreated:   time.Now(),
		FCMKeyUpdated: time.Now(),
	}
	db.Create(&u)
	return c.JSON(http.StatusCreated, u)
}
