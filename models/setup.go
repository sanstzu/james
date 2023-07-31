package models

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var client *firestore.Client

func InitializeFirebase() {
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal(err)
	}

	newClient, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	client = newClient
}

func GetClient() *firestore.Client {
	return client
}
