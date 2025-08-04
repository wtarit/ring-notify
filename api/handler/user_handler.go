package handler

import (
	"api/configs"
	"api/models"
	"api/service"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	UserHandler struct {
		service *service.UserService
	}
	CreateUserRequest struct {
		FcmToken string `json:"fcmToken" validate:"required" example:"fcm-token"`
	}
	CreateUserResponse struct {
		ID            string `json:"id" example:"00000000-0000-0000-0000-000000000000"`
		APIKey        string `json:"apiKey" example:"00000000-0000-0000-0000-000000000000"`
		FCMKey        string `json:"fcmKey" example:"fcm-token-example"`
		UserCreated   string `json:"userCreated" example:"2025-01-01T00:00:00Z"`
		FCMKeyUpdated string `json:"fcmKeyUpdated" example:"2025-01-01T00:00:00Z"`
	}
	CustomValidator struct {
		Validator *validator.Validate
	}
)

func NewUserHandler() *UserHandler {
	return &UserHandler{service: service.NewUserService()}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

// CreateUser godoc
//
//	@Summary		Create a new user
//	@Description	Create a new user with FCM token and get API key
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateUserRequest	true	"User creation request"
//	@Success		201		{object}	CreateUserResponse
//	@Failure		400		{string}	string	"Bad Request"
//	@Router			/user/create [post]
func (h *UserHandler) CreateUser(c echo.Context) error {
	var reqBody CreateUserRequest
	err := c.Bind(&reqBody)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	if err = c.Validate(reqBody); err != nil {
		return err
	}
	u := h.service.CreateUser(reqBody.FcmToken)
	return c.JSON(http.StatusCreated, u)
}

func (h *UserHandler) RefreshToken(c echo.Context) {
	apikey := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]
	db := configs.DB()
	var user models.User
	db.First(&user, "api_key = ?", apikey)
}
