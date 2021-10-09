package firebase

import (
	"canny/pkg/log"
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"go.uber.org/zap"
	"google.golang.org/api/option"
)

var (
	client *firestore.Client
	ctx    context.Context
)

func Initialise() {
	// Use a service account
	ctx = context.Background()
	sa := option.WithCredentialsFile("config/serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Logger.Fatal("Failed to initialise firebase ", zap.Error(err))
	}

	client, err = app.Firestore(ctx)
	if err != nil {
		log.Logger.Fatal("Failed to initialise firebase client ", zap.Error(err))
	}
	//defer client.Close()
}

func Add(collection string, data interface{}) {
	_, _, err := client.Collection(collection).Add(ctx, data)
	if err != nil {
		log.Logger.Fatal("Failed adding data to firebase : %v", err)
	}
}
