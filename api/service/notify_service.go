package service

import (
	"api/configs"
	"context"
	"errors"
	"log"

	"firebase.google.com/go/v4/messaging"
)

type NotifyService struct{}

// NotifyResult contains the results of sending notifications to multiple devices
type NotifyResult struct {
	SuccessCount int      `json:"successCount"`
	FailureCount int      `json:"failureCount"`
	FailedTokens []string `json:"failedTokens,omitempty"`
}

func NewNotifyService() *NotifyService {
	return &NotifyService{}
}

// NotifyMultiple sends notifications to multiple FCM tokens
func (s *NotifyService) NotifyMultiple(fcmTokens []string, notificationText string) (*NotifyResult, error) {
	if len(fcmTokens) == 0 {
		return nil, errors.New("no FCM tokens provided")
	}

	ctx := context.Background()
	client, err := configs.App.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	result := &NotifyResult{
		SuccessCount: 0,
		FailureCount: 0,
		FailedTokens: make([]string, 0),
	}

	deviceService := NewDeviceService()

	// Send to all tokens
	for _, token := range fcmTokens {
		message := &messaging.Message{
			Data: map[string]string{
				"text": notificationText,
			},
			Token: token,
			Android: &messaging.AndroidConfig{
				Priority: "high",
			},
		}

		_, err = client.Send(ctx, message)
		if err != nil {
			log.Printf("FCM error for token %s: %v\n", token, err)
			result.FailureCount++
			result.FailedTokens = append(result.FailedTokens, token)

			// Mark device as inactive if token is invalid
			go deviceService.MarkDeviceInactive(token)
		} else {
			result.SuccessCount++

			// Update last active
			go deviceService.UpdateDeviceLastActive(token)
		}
	}

	return result, nil
}
