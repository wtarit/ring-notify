package models

import "time"

// RegisterDeviceRequest represents a request to register a new device
type RegisterDeviceRequest struct {
	FCMToken   string `json:"fcmToken" validate:"required" example:"fcm-token-example"`
	DeviceName string `json:"deviceName" validate:"required" example:"My Phone"`
	DeviceType string `json:"deviceType" validate:"required" example:"android"`
}

// UpdateDeviceRequest represents a request to update device information
type UpdateDeviceRequest struct {
	DeviceName *string `json:"deviceName,omitempty" example:"My Updated Phone"`
	FCMToken   *string `json:"fcmToken,omitempty" example:"new-fcm-token"`
}

// DeviceResponse represents a device in API responses
type DeviceResponse struct {
	ID           string    `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	DeviceName   string    `json:"deviceName" example:"My Phone"`
	DeviceType   string    `json:"deviceType" example:"android"`
	RegisteredAt time.Time `json:"registeredAt" example:"2023-01-01T00:00:00Z"`
	LastActive   time.Time `json:"lastActive" example:"2023-01-01T00:00:00Z"`
	IsActive     bool      `json:"isActive" example:"true"`
}

// DeviceListResponse represents a list of devices
type DeviceListResponse struct {
	Devices []DeviceResponse `json:"devices"`
}
