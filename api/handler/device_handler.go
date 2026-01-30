package handler

import (
	"api/ctxutil"
	"api/models"
	"api/service"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// DeviceHandler handles device-related HTTP requests
type DeviceHandler struct {
	service *service.DeviceService
}

// NewDeviceHandler creates a new DeviceHandler
func NewDeviceHandler() *DeviceHandler {
	return &DeviceHandler{
		service: service.NewDeviceService(),
	}
}

// Register godoc
//
//	@Summary		Register a new device
//	@Description	Register a device to receive notifications
//	@Tags			devices
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.RegisterDeviceRequest	true	"Device registration details"
//	@Success		201		{object}	models.DeviceResponse
//	@Failure		400		{object}	models.BadRequestResponse
//	@Failure		401		{object}	models.BadRequestResponse
//	@Failure		500		{object}	models.BadRequestResponse
//	@Security		BearerAuth
//	@Router			/devices [post]
func (h *DeviceHandler) Register(c echo.Context) error {
	supabaseUserID := ctxutil.GetSupabaseUserID(c)

	var req models.RegisterDeviceRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.NewErrorResponse("Invalid request"))
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.NewErrorResponse("Validation failed"))
	}

	device, err := h.service.RegisterDevice(supabaseUserID, req.FCMToken, req.DeviceName, req.DeviceType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.NewErrorResponse("Failed to register device"))
	}

	response := models.DeviceResponse{
		ID:           device.ID.String(),
		DeviceName:   device.DeviceName,
		DeviceType:   device.DeviceType,
		RegisteredAt: device.RegisteredAt,
		LastActive:   device.LastActive,
		IsActive:     device.IsActive,
	}

	return c.JSON(http.StatusCreated, response)
}

// List godoc
//
//	@Summary		List all devices
//	@Description	Get a list of all devices registered for the authenticated user
//	@Tags			devices
//	@Produce		json
//	@Success		200	{object}	models.DeviceListResponse
//	@Failure		401	{object}	models.BadRequestResponse
//	@Failure		500	{object}	models.BadRequestResponse
//	@Security		BearerAuth
//	@Router			/devices [get]
func (h *DeviceHandler) List(c echo.Context) error {
	supabaseUserID := ctxutil.GetSupabaseUserID(c)

	devices, err := h.service.ListDevices(supabaseUserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.NewErrorResponse("Failed to list devices"))
	}

	deviceResponses := make([]models.DeviceResponse, len(devices))
	for i, device := range devices {
		deviceResponses[i] = models.DeviceResponse{
			ID:           device.ID.String(),
			DeviceName:   device.DeviceName,
			DeviceType:   device.DeviceType,
			RegisteredAt: device.RegisteredAt,
			LastActive:   device.LastActive,
			IsActive:     device.IsActive,
		}
	}

	return c.JSON(http.StatusOK, models.DeviceListResponse{Devices: deviceResponses})
}

// Update godoc
//
//	@Summary		Update device information
//	@Description	Update device name or FCM token
//	@Tags			devices
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string							true	"Device ID"
//	@Param			request	body		models.UpdateDeviceRequest		true	"Device update details"
//	@Success		200		{object}	models.DeviceResponse
//	@Failure		400		{object}	models.BadRequestResponse
//	@Failure		401		{object}	models.BadRequestResponse
//	@Failure		404		{object}	models.BadRequestResponse
//	@Failure		500		{object}	models.BadRequestResponse
//	@Security		BearerAuth
//	@Router			/devices/{id} [patch]
func (h *DeviceHandler) Update(c echo.Context) error {
	supabaseUserID := ctxutil.GetSupabaseUserID(c)

	deviceID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.NewErrorResponse("Invalid device ID"))
	}

	var req models.UpdateDeviceRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.NewErrorResponse("Invalid request"))
	}

	device, err := h.service.UpdateDevice(supabaseUserID, deviceID, req.DeviceName, req.FCMToken)
	if err != nil {
		if err.Error() == "device not found" {
			return c.JSON(http.StatusNotFound, models.NewErrorResponse("Device not found"))
		}
		return c.JSON(http.StatusInternalServerError, models.NewErrorResponse("Failed to update device"))
	}

	response := models.DeviceResponse{
		ID:           device.ID.String(),
		DeviceName:   device.DeviceName,
		DeviceType:   device.DeviceType,
		RegisteredAt: device.RegisteredAt,
		LastActive:   device.LastActive,
		IsActive:     device.IsActive,
	}

	return c.JSON(http.StatusOK, response)
}

// Remove godoc
//
//	@Summary		Remove a device
//	@Description	Remove a device from the user's account
//	@Tags			devices
//	@Produce		json
//	@Param			id	path		string	true	"Device ID"
//	@Success		200	{object}	models.SuccessResponse
//	@Failure		400	{object}	models.BadRequestResponse
//	@Failure		401	{object}	models.BadRequestResponse
//	@Failure		404	{object}	models.BadRequestResponse
//	@Failure		500	{object}	models.BadRequestResponse
//	@Security		BearerAuth
//	@Router			/devices/{id} [delete]
func (h *DeviceHandler) Remove(c echo.Context) error {
	supabaseUserID := ctxutil.GetSupabaseUserID(c)

	deviceID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.NewErrorResponse("Invalid device ID"))
	}

	if err := h.service.RemoveDevice(supabaseUserID, deviceID); err != nil {
		if err.Error() == "device not found" {
			return c.JSON(http.StatusNotFound, models.NewErrorResponse("Device not found"))
		}
		return c.JSON(http.StatusInternalServerError, models.NewErrorResponse("Failed to remove device"))
	}

	return c.JSON(http.StatusOK, models.SuccessResponse{Message: "Device removed successfully"})
}
