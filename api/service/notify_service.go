package service

import (
	"api/configs"
	"context"
	"errors"
	"log"

	"firebase.google.com/go/v4/messaging"
)

type NotifyService struct{}

func NewNotifyService() *NotifyService {
	return &NotifyService{}
}

func (s *NotifyService) Notify(fcmKey string, notificationText string) error {
	// Obtain a messaging.Client from the App.
	ctx := context.Background()
	client, err := configs.App.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: map[string]string{
			"text": notificationText,
		},
		Token: fcmKey,
		Android: &messaging.AndroidConfig{
			Priority: "high",
		},
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	_, err = client.Send(ctx, message)
	if err != nil {
		log.Printf("FCM error: %v\n", err)
		return errors.New("token no longer valid")
	}
	return nil
}
