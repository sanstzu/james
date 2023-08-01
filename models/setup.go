package models

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/sanstzu/james/consts"
	"google.golang.org/api/option"
)

var client *firestore.Client

func InitializeFirebase() {
	opt := option.WithCredentialsFile(consts.ENV("FIREBASE_SERVICEACCOUNT_KEY_PATH"))
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
