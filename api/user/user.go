package user

import (
	"api/configs"
	"api/models"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type (
	CreateUserRequest struct {
		FcmToken string `json:"fcmToken" validate:"required"`
	}
	CustomValidator struct {
		Validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func CreateUser(c echo.Context) error {
	var reqBody CreateUserRequest
	err := c.Bind(&reqBody)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err = c.Validate(reqBody); err != nil {
		return err
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

func RefreshToken(c echo.Context) {
	apikey := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]
	db := configs.DB()
	var user models.User
	db.First(&user, "api_key = ?", apikey)
}
