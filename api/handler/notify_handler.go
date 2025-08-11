package handler

import (
	"api/ctxutil"
	"api/models"
	"api/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type NotifyHandler struct {
	service *service.NotifyService
}

func NewNotifyHandler() *NotifyHandler {
	return &NotifyHandler{
		service: service.NewNotifyService(),
	}
}

// Call godoc
//
//	@Summary		Send FCM notification call
//	@Description	Send a Firebase Cloud Messaging notification to trigger a call
//	@Tags			notify
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"Bearer token (API Key)"	default(Bearer your-api-key-here)
//	@Param			request			body		models.CallRequest	true	"Call request payload"
//	@Success		200				{object}	models.SuccessResponse
//	@Failure		400				{object}	models.BadRequestResponse
//	@Failure		500				{object}	models.NotifyErrorResponse
//	@Security		BearerAuth
//	@Router			/notify/call [post]
func (h *NotifyHandler) Call(c echo.Context) error {
	var callRequest models.CallRequest
	err := c.Bind(&callRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.NewErrorResponse("Bad Request"))
	}
	user := ctxutil.GetUser(c)

	registrationToken := user.FCMKey

	err = h.service.Notify(registrationToken, callRequest.Text)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.NotifyErrorResponse{
			Reason: "Failed to send call",
		})
	}
	return c.JSON(http.StatusOK, models.NewSuccessResponse("Called"))
}
