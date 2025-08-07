package models

type CreateUserRequest struct {
	FcmToken string `json:"fcmToken" validate:"required" example:"fcm-token"`
}

type CreateUserResponse struct {
	ID            string `json:"id" example:"00000000-0000-0000-0000-000000000000"`
	APIKey        string `json:"apiKey" example:"00000000-0000-0000-0000-000000000000"`
	FCMKey        string `json:"fcmKey" example:"fcm-token-example"`
	UserCreated   string `json:"userCreated" example:"2025-01-01T00:00:00Z"`
	FCMKeyUpdated string `json:"fcmKeyUpdated" example:"2025-01-01T00:00:00Z"`
}
