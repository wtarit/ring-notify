package handler

import (
	"api/ctxutil"
	"api/models"
	"api/service"
	"net/http"

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
//	@Description	Create a new anonymous user with FCM token and get API key
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.CreateUserRequest	true	"User creation request"
//	@Success		201		{object}	models.CreateUserResponse
//	@Failure		400		{object}	models.BadRequestResponse
//	@Router			/users [post]
func (h *UserHandler) CreateUser(c echo.Context) error {
	var reqBody models.CreateUserRequest
	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, models.NewErrorResponse(err.Error()))
	}
	if err := c.Validate(reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, models.NewErrorResponse("Validation failed"))
	}

	result := h.service.CreateUser(reqBody.FcmToken)
	if result == nil {
		return c.JSON(http.StatusInternalServerError, models.NewErrorResponse("Failed to create user"))
	}

	resp := models.CreateUserResponse{
		ID:     result.User.ID.String(),
		APIKey: result.APIKey,
	}
	return c.JSON(http.StatusCreated, resp)
}

// RegenerateAPIKey godoc
//
//	@Summary      Regenerate API key
//	@Description  Regenerates the API key for the authenticated user (creates new, deactivates old)
//	@Tags         users
//	@Produce      json
//	@Success      200  {object}  models.CreateUserResponse
//	@Failure      400  {object}  models.BadRequestResponse
//	@Failure      401  {object}  models.BadRequestResponse
//	@Security     BearerAuth
//	@Router       /users/api-key [post]
func (h *UserHandler) RegenerateAPIKey(c echo.Context) error {
	user := ctxutil.GetUser(c)
	if user == nil {
		return c.JSON(http.StatusUnauthorized, models.NewErrorResponse("Unauthorized"))
	}

	newAPIKey, err := h.service.RegenerateAPIKey(user.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.NewErrorResponse("Failed to regenerate API key"))
	}

	resp := models.CreateUserResponse{
		ID:     user.ID.String(),
		APIKey: newAPIKey,
	}
	return c.JSON(http.StatusOK, resp)
}
