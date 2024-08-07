package notify

import (
	"api/configs"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

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

	// This registration token comes from the client FCM SDKs.
	registrationToken := os.Getenv("TMP_CLIENT")

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
