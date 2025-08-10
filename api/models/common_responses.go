package models

// Common response models for consistent API responses

type SuccessResponse struct {
	Message string `json:"message"`
}

type BadRequestResponse struct {
	Error string `json:"error" example:"Bad Request"`
}

type ValidationErrorResponse struct {
	Error   string            `json:"error"`
	Details map[string]string `json:"details,omitempty"`
}

// Helper functions for common responses
func NewSuccessResponse(message string) *SuccessResponse {
	return &SuccessResponse{Message: message}
}

func NewErrorResponse(error string) *BadRequestResponse {
	return &BadRequestResponse{Error: error}
}

func NewValidationErrorResponse(error string, details map[string]string) *ValidationErrorResponse {
	return &ValidationErrorResponse{
		Error:   error,
		Details: details,
	}
}
