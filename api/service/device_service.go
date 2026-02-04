package service

import (
	"api/configs"
	"api/models"
	"errors"
	"time"

	"github.com/google/uuid"
)

// DeviceService handles device-related business logic
type DeviceService struct{}

// NewDeviceService creates a new DeviceService
func NewDeviceService() *DeviceService {
	return &DeviceService{}
}

// RegisterDevice registers a new device for a user
func (s *DeviceService) RegisterDevice(supabaseUserID uuid.UUID, fcmToken, deviceName, deviceType string) (*models.Device, error) {
	db := configs.DB()

	// Check if device already exists for this user with this FCM token
	var existing models.Device
	if err := db.First(&existing, "supabase_user_id = ? AND fcm_token = ?", supabaseUserID, fcmToken).Error; err == nil {
		// Device exists, update it
		existing.DeviceName = deviceName
		existing.DeviceType = deviceType
		existing.IsActive = true
		existing.LastActive = time.Now()
		if err := db.Save(&existing).Error; err != nil {
			return nil, err
		}
		return &existing, nil
	}

	// Create new device
	device := models.Device{
		ID:           uuid.New(),
		UserID:       supabaseUserID,
		FCMToken:     fcmToken,
		DeviceName:   deviceName,
		DeviceType:   deviceType,
		RegisteredAt: time.Now(),
		LastActive:   time.Now(),
		IsActive:     true,
	}

	if err := db.Create(&device).Error; err != nil {
		return nil, err
	}

	return &device, nil
}

// ListDevices returns all devices for a user
func (s *DeviceService) ListDevices(supabaseUserID uuid.UUID) ([]models.Device, error) {
	db := configs.DB()
	var devices []models.Device
	err := db.Where("user_id = ?", supabaseUserID).Order("registered_at DESC").Find(&devices).Error
	return devices, err
}

// RemoveDevice removes a device (soft delete)
func (s *DeviceService) RemoveDevice(supabaseUserID uuid.UUID, deviceID uuid.UUID) error {
	db := configs.DB()
	result := db.Where("id = ? AND supabase_user_id = ?", deviceID, supabaseUserID).Delete(&models.Device{})
	if result.RowsAffected == 0 {
		return errors.New("device not found")
	}
	return result.Error
}

// UpdateDevice updates device information
func (s *DeviceService) UpdateDevice(supabaseUserID uuid.UUID, deviceID uuid.UUID, deviceName *string, fcmToken *string) (*models.Device, error) {
	db := configs.DB()
	var device models.Device

	if err := db.First(&device, "id = ? AND supabase_user_id = ?", deviceID, supabaseUserID).Error; err != nil {
		return nil, errors.New("device not found")
	}

	if deviceName != nil {
		device.DeviceName = *deviceName
	}

	if fcmToken != nil {
		device.FCMToken = *fcmToken
	}

	device.LastActive = time.Now()

	if err := db.Save(&device).Error; err != nil {
		return nil, err
	}

	return &device, nil
}

// MarkDeviceInactive marks a device as inactive (called when FCM fails)
func (s *DeviceService) MarkDeviceInactive(fcmToken string) error {
	db := configs.DB()
	return db.Model(&models.Device{}).
		Where("fcm_token = ?", fcmToken).
		Update("is_active", false).Error
}

// UpdateDeviceLastActive updates the last active timestamp
func (s *DeviceService) UpdateDeviceLastActive(fcmToken string) error {
	db := configs.DB()
	return db.Model(&models.Device{}).
		Where("fcm_token = ?", fcmToken).
		Update("last_active", time.Now()).Error
}
