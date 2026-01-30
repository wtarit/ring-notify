package models

type CreateUserRequest struct {
	FcmToken string `json:"fcmToken" validate:"required" example:"fcm-token"`
}

type CreateUserResponse struct {
	ID     string `json:"id" example:"00000000-0000-0000-0000-000000000000"`
	APIKey string `json:"apiKey" example:"00000000-0000-0000-0000-000000000000"`
}
