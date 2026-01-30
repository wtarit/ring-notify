package handler

import (
	"api/configs"
	"api/ctxutil"
	"api/models"
	"api/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type NotifyHandler struct {
	service       *service.NotifyService
	deviceService *service.DeviceService
}

func NewNotifyHandler() *NotifyHandler {
	return &NotifyHandler{
		service:       service.NewNotifyService(),
		deviceService: service.NewDeviceService(),
	}
}

// NotifyMultipleResponse represents the response for multi-device notifications
type NotifyMultipleResponse struct {
	Message      string `json:"message"`
	SuccessCount int    `json:"successCount"`
	FailureCount int    `json:"failureCount"`
}

// Call godoc
//
//	@Summary		Send FCM notification call
//	@Description	Send a Firebase Cloud Messaging notification to trigger a call on all registered devices
//	@Tags			notify
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string				true	"Bearer token (API Key)"	default(Bearer your-api-key-here)
//	@Param			request			body		models.CallRequest	true	"Call request payload"
//	@Success		200				{object}	NotifyMultipleResponse
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

	if err := c.Validate(&callRequest); err != nil {
		return c.JSON(http.StatusBadRequest, models.NewErrorResponse("Invalid request"))
	}

	supabaseUserID := ctxutil.GetSupabaseUserID(c)

	// Fetch all active devices for the user
	db := configs.DB()
	var devices []models.Device
	if err := db.Where("user_id = ? AND is_active = ?", supabaseUserID, true).Find(&devices).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, &models.NotifyErrorResponse{
			Reason: "Failed to fetch devices",
		})
	}

	if len(devices) == 0 {
		return c.JSON(http.StatusBadRequest, models.NewErrorResponse("No active devices found"))
	}

	// Collect FCM tokens
	fcmTokens := make([]string, len(devices))
	for i, device := range devices {
		fcmTokens[i] = device.FCMToken
	}

	// Send to all devices
	result, err := h.service.NotifyMultiple(fcmTokens, callRequest.Text)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &models.NotifyErrorResponse{
			Reason: "Failed to send notifications",
		})
	}

	return c.JSON(http.StatusOK, NotifyMultipleResponse{
		Message:      "Notifications sent",
		SuccessCount: result.SuccessCount,
		FailureCount: result.FailureCount,
	})
}
