package notify

import (
	"api/configs"
	"api/models"
	"context"
	"log"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/messaging"
	"github.com/labstack/echo/v4"
)

type CallRequest struct {
	Text string `json:"text" validate:"required" example:"Notification from ESP32"`
}

type ErrorResponse struct {
	Reason string `json:"reason" example:"Token no longer valid"`
}

// Call godoc
//
//	@Summary		Send FCM notification call
//	@Description	Send a Firebase Cloud Messaging notification to trigger a call
//	@Tags			notify
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string		true	"Bearer token (API Key)"	default(Bearer your-api-key-here)
//	@Param			request			body		CallRequest	true	"Call request payload"
//	@Success		200				{string}	string		"Called"
//	@Failure		400				{string}	string		"Bad Request"
//	@Failure		403				{object}	ErrorResponse
//	@Security		BearerAuth
//	@Router			/notify/call [post]
func Call(c echo.Context) error {
	var callRequest CallRequest
	err := c.Bind(&callRequest)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}
	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := configs.App.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}
	apikey := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]
	db := configs.DB()
	// This registration token comes from the client FCM SDKs.
	var user models.User
	db.First(&user, "api_key = ?", apikey)
	registrationToken := user.FCMKey

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: map[string]string{
			"text": callRequest.Text,
		},
		Token: registrationToken,
		Android: &messaging.AndroidConfig{
			Priority: "high",
		},
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	_, err = client.Send(ctx, message)
	if err != nil {
		log.Printf("FCM error: %v\n", err)
		return c.JSON(http.StatusForbidden, &ErrorResponse{
			Reason: "Token no longer valid",
		})
	}
	return c.String(http.StatusOK, "Called")
}
