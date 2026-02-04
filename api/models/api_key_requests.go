package models

import "time"

// CreateAPIKeyRequest represents a request to create a new API key
type CreateAPIKeyRequest struct {
	Name      string     `json:"name" validate:"required" example:"ESP32 Doorbell"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty" example:"2024-01-01T00:00:00Z"`
}

// APIKeyResponse represents an API key in responses
type APIKeyResponse struct {
	ID         string     `json:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Key        string     `json:"key" example:"550e8400-e29b-41d4-a716-446655440002"`
	Name       string     `json:"name" example:"ESP32 Doorbell"`
	CreatedAt  time.Time  `json:"createdAt" example:"2023-01-01T00:00:00Z"`
	ExpiresAt  *time.Time `json:"expiresAt,omitempty" example:"2024-01-01T00:00:00Z"`
	LastUsedAt *time.Time `json:"lastUsedAt,omitempty" example:"2023-01-01T00:00:00Z"`
	IsActive   bool       `json:"isActive" example:"true"`
}

// APIKeyListResponse represents a list of API keys (with masked keys)
type APIKeyListResponse struct {
	APIKeys []APIKeyResponse `json:"apiKeys"`
}
