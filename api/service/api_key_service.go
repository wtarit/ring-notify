package service

import (
	"api/configs"
	"api/models"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// APIKeyService handles API key-related business logic
type APIKeyService struct{}

// NewAPIKeyService creates a new APIKeyService
func NewAPIKeyService() *APIKeyService {
	return &APIKeyService{}
}

// CreateAPIKey creates a new API key for a user
func (s *APIKeyService) CreateAPIKey(supabaseUserID uuid.UUID, name string, expiresAt *time.Time) (*models.APIKey, error) {
	db := configs.DB()

	// If no expiry provided, use default (1 year)
	if expiresAt == nil {
		defaultDays := 365
		if envDays := os.Getenv("API_KEY_DEFAULT_EXPIRY_DAYS"); envDays != "" {
			if days, err := strconv.Atoi(envDays); err == nil {
				defaultDays = days
			}
		}
		defaultExpiry := time.Now().AddDate(0, 0, defaultDays)
		expiresAt = &defaultExpiry
	}

	apiKey := models.APIKey{
		ID:        uuid.New(),
		UserID:    supabaseUserID,
		Key:       uuid.NewString(),
		Name:      name,
		CreatedAt: time.Now(),
		ExpiresAt: expiresAt,
		IsActive:  true,
	}

	if err := db.Create(&apiKey).Error; err != nil {
		return nil, err
	}

	return &apiKey, nil
}

// ListAPIKeys returns all API keys for a user (with masked keys)
func (s *APIKeyService) ListAPIKeys(supabaseUserID uuid.UUID) ([]models.APIKey, error) {
	db := configs.DB()
	var apiKeys []models.APIKey
	err := db.Where("user_id = ?", supabaseUserID).Order("created_at DESC").Find(&apiKeys).Error

	// Mask the keys in list view for security
	for i := range apiKeys {
		if len(apiKeys[i].Key) > 8 {
			apiKeys[i].Key = apiKeys[i].Key[:4] + "••••" + apiKeys[i].Key[len(apiKeys[i].Key)-4:]
		} else {
			apiKeys[i].Key = "••••••••"
		}
	}

	return apiKeys, err
}

// RevokeAPIKey marks an API key as inactive
func (s *APIKeyService) RevokeAPIKey(supabaseUserID uuid.UUID, apiKeyID uuid.UUID) error {
	db := configs.DB()
	result := db.Model(&models.APIKey{}).
		Where("id = ? AND supabase_user_id = ?", apiKeyID, supabaseUserID).
		Update("is_active", false)

	if result.RowsAffected == 0 {
		return errors.New("API key not found")
	}

	return result.Error
}

// GetAPIKeyByKey retrieves an API key by its key string (for auth middleware)
func (s *APIKeyService) GetAPIKeyByKey(keyStr string) (*models.APIKey, error) {
	db := configs.DB()
	var apiKey models.APIKey
	err := db.First(&apiKey, "key = ? AND is_active = ?", keyStr, true).Error
	if err != nil {
		return nil, err
	}

	// Check expiry
	if apiKey.ExpiresAt != nil && apiKey.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("API key expired")
	}

	return &apiKey, nil
}
