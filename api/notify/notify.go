package notify

import (
	"api/configs"
	"api/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/messaging"
	"github.com/labstack/echo/v4"
)

type CallRequest struct {
	Text string `json:"text"`
}

func Call(c echo.Context) error {
	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := configs.App.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}
	apikey := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]
	fmt.Printf("Authorization header: %s\n", apikey)
	db := configs.DB()
	// This registration token comes from the client FCM SDKs.
	var user models.User
	db.First(&user, "api_key = ?", apikey)
	registrationToken := user.FCMKey

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: map[string]string{
			"text": "test",
		},
		Token: registrationToken,
		Android: &messaging.AndroidConfig{
			Priority: "high",
		},
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
	return c.String(http.StatusOK, "Called")
}
