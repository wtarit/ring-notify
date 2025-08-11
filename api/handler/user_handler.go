package handler

import (
	"api/models"
	"api/service"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type (
	UserHandler struct {
		service *service.UserService
	}
)

func NewUserHandler() *UserHandler {
	return &UserHandler{service: service.NewUserService()}
}

// CreateUser godoc
//
//	@Summary		Create a new user
//	@Description	Create a new user with FCM token and get API key
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.CreateUserRequest	true	"User creation request"
//	@Success		201		{object}	models.CreateUserResponse
//	@Failure		400		{object}	models.BadRequestResponse
//	@Router			/user [post]
func (h *UserHandler) CreateUser(c echo.Context) error {
	var reqBody models.CreateUserRequest
	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, models.NewErrorResponse(err.Error()))
	}
	if err := c.Validate(reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, models.NewErrorResponse("Validation failed"))
	}
	u := h.service.CreateUser(reqBody.FcmToken)
	resp := models.CreateUserResponse{
		ID:            u.ID.String(),
		APIKey:        u.APIKey,
		FCMKey:        u.FCMKey,
		UserCreated:   u.UserCreated.Format(time.RFC3339),
		FCMKeyUpdated: u.FCMKeyUpdated.Format(time.RFC3339),
	}
	return c.JSON(http.StatusCreated, resp)
}
