package handler

import (
	"api/models"
	"api/service"
	"net/http"
	"strings"

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
//	@Param			Authorization	header		string		true	"Bearer token (API Key)"	default(Bearer your-api-key-here)
//	@Param			request			body		CallRequest	true	"Call request payload"
//	@Success		200				{string}	string		"Called"
//	@Failure		400				{string}	string		"Bad Request"
//	@Failure		403				{object}	ErrorResponse
//	@Security		BearerAuth
//	@Router			/notify/call [post]
func (h *NotifyHandler) Call(c echo.Context) error {
	var callRequest models.CallRequest
	err := c.Bind(&callRequest)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	apiKey := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]
	err = h.service.Notify(apiKey, callRequest.Text)
	if err != nil {
		return c.JSON(http.StatusForbidden, &models.ErrorResponse{
			Reason: "Error",
		})
	}
	return c.String(http.StatusOK, "Called")
}
