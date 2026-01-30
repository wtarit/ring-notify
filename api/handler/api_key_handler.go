package handler

import (
	"api/ctxutil"
	"api/models"
	"api/service"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// APIKeyHandler handles API key-related HTTP requests
type APIKeyHandler struct {
	service *service.APIKeyService
}

// NewAPIKeyHandler creates a new APIKeyHandler
func NewAPIKeyHandler() *APIKeyHandler {
	return &APIKeyHandler{
		service: service.NewAPIKeyService(),
	}
}

// Create godoc
//
//	@Summary		Create a new API key
//	@Description	Create a new API key with optional expiry date
//	@Tags			api-keys
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.CreateAPIKeyRequest	true	"API key creation details"
//	@Success		201		{object}	models.APIKeyResponse
//	@Failure		400		{object}	models.BadRequestResponse
//	@Failure		401		{object}	models.BadRequestResponse
//	@Failure		500		{object}	models.BadRequestResponse
//	@Security		BearerAuth
//	@Router			/api-keys [post]
func (h *APIKeyHandler) Create(c echo.Context) error {
	supabaseUserID := ctxutil.GetSupabaseUserID(c)

	var req models.CreateAPIKeyRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.NewErrorResponse("Invalid request"))
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.NewErrorResponse("Validation failed"))
	}

	apiKey, err := h.service.CreateAPIKey(supabaseUserID, req.Name, req.ExpiresAt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.NewErrorResponse("Failed to create API key"))
	}

	response := models.APIKeyResponse{
		ID:         apiKey.ID.String(),
		Key:        apiKey.Key, // Show full key only once on creation
		Name:       apiKey.Name,
		CreatedAt:  apiKey.CreatedAt,
		ExpiresAt:  apiKey.ExpiresAt,
		LastUsedAt: apiKey.LastUsedAt,
		IsActive:   apiKey.IsActive,
	}

	return c.JSON(http.StatusCreated, response)
}

// List godoc
//
//	@Summary		List all API keys
//	@Description	Get a list of all API keys for the authenticated user (keys are masked)
//	@Tags			api-keys
//	@Produce		json
//	@Success		200	{object}	models.APIKeyListResponse
//	@Failure		401	{object}	models.BadRequestResponse
//	@Failure		500	{object}	models.BadRequestResponse
//	@Security		BearerAuth
//	@Router			/api-keys [get]
func (h *APIKeyHandler) List(c echo.Context) error {
	supabaseUserID := ctxutil.GetSupabaseUserID(c)

	apiKeys, err := h.service.ListAPIKeys(supabaseUserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.NewErrorResponse("Failed to list API keys"))
	}

	apiKeyResponses := make([]models.APIKeyResponse, len(apiKeys))
	for i, apiKey := range apiKeys {
		apiKeyResponses[i] = models.APIKeyResponse{
			ID:         apiKey.ID.String(),
			Key:        apiKey.Key, // Already masked by service
			Name:       apiKey.Name,
			CreatedAt:  apiKey.CreatedAt,
			ExpiresAt:  apiKey.ExpiresAt,
			LastUsedAt: apiKey.LastUsedAt,
			IsActive:   apiKey.IsActive,
		}
	}

	return c.JSON(http.StatusOK, models.APIKeyListResponse{APIKeys: apiKeyResponses})
}

// Revoke godoc
//
//	@Summary		Revoke an API key
//	@Description	Mark an API key as inactive
//	@Tags			api-keys
//	@Produce		json
//	@Param			id	path		string	true	"API Key ID"
//	@Success		200	{object}	models.SuccessResponse
//	@Failure		400	{object}	models.BadRequestResponse
//	@Failure		401	{object}	models.BadRequestResponse
//	@Failure		404	{object}	models.BadRequestResponse
//	@Failure		500	{object}	models.BadRequestResponse
//	@Security		BearerAuth
//	@Router			/api-keys/{id} [delete]
func (h *APIKeyHandler) Revoke(c echo.Context) error {
	supabaseUserID := ctxutil.GetSupabaseUserID(c)

	apiKeyID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.NewErrorResponse("Invalid API key ID"))
	}

	if err := h.service.RevokeAPIKey(supabaseUserID, apiKeyID); err != nil {
		if err.Error() == "API key not found" {
			return c.JSON(http.StatusNotFound, models.NewErrorResponse("API key not found"))
		}
		return c.JSON(http.StatusInternalServerError, models.NewErrorResponse("Failed to revoke API key"))
	}

	return c.JSON(http.StatusOK, models.SuccessResponse{Message: "API key revoked successfully"})
}
