package models

type CallRequest struct {
	Text string `json:"text" validate:"required" example:"Notification from ESP32"`
}

type NotifyErrorResponse struct {
	Reason string `json:"reason" example:"Token no longer valid"`
}
