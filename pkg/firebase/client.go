package firebase

import (
	err2 "canny/pkg/err"
	"canny/pkg/log"
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
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
}

func Shutdown() {
	defer func(client *firestore.Client) {
		err := client.Close()
		if err != nil {
			log.Logger.Fatal("Failed to close firebase connection : %v", err)
		}
	}(client)
}

func GetAll(collection string) []map[string]interface{} {
	items := make([]map[string]interface{}, 0)
	iter := client.Collection(collection).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Logger.Fatal("Failed getting alerts from firebase : %v", err)
		}
		data := doc.Data()
		data["createdAt"] = doc.CreateTime
		data["id"] = doc.Ref.ID
		items = append(items, data)
	}
	return items
}

func Get(collection string, id string) map[string]interface{} {
	var data map[string]interface{}
	item, err := client.Collection(collection).Doc(id).Get(ctx)
	if err != nil {
		log.Logger.Fatal("Failed getting alert from firebase : %v", err)
	} else {
		data = item.Data()
		data["createdAt"] = item.CreateTime
		data["id"] = id
	}
	return data
}

func Add(collection string, data interface{}) map[string]interface{} {
	documentRef, _, err := client.Collection(collection).Add(ctx, data)
	if err != nil {
		log.Logger.Fatal("Failed adding alert to firebase : %v", err)
	}
	savedData := Get(collection, documentRef.ID)
	return savedData
}

func Delete(collection string, id string) (error, int) {
	_, err := client.Collection(collection).Doc(id).Delete(ctx)
	if err != nil {
		log.Logger.Fatal("Failed deleting alert from firebase : %v", err)
		if len(Get(collection, id)) == 0 {
			return nil, err2.NotFound
		}
		return err, 0
	}
	return nil, 0
}
