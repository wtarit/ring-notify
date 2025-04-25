package configs

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var App *firebase.App

func InitFirebase() {
	var err error
	opt := option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_CONFIG")))

	App, err = firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	fmt.Println("Done configuring firebase.")
}
