package firebase

import (
	"canny/model"
	err2 "canny/pkg/err"
	"canny/pkg/log"
	"canny/pkg/utils"
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

func GetAll(collection string) []model.CreateAlertResponseData {
	items := make([]model.CreateAlertResponseData, 0)
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
		convData := model.CreateAlertResponseData{
			Id:         doc.Ref.ID,
			Name:       data["Name"].(string),
			Property:   data["Property"].(string),
			Operator:   data["Operator"].(string),
			Value:      data["Value"].(float64),
			Percentage: data["Percentage"].(int64),
			Currency:   data["Currency"].(string),
			CreatedAt:  doc.CreateTime,
		}
		items = append(items, convData)
	}
	return items
}

func Get(collection string, id string) (map[string]interface{}, error, int) {
	var data map[string]interface{}
	item, err := client.Collection(collection).Doc(id).Get(ctx)
	if err != nil && utils.Is404Error(err.Error()) {
		return nil, nil, err2.NotFound
	} else if err != nil {
		log.Logger.Fatal("Failed getting alert from firebase : %v", err)
		return nil, err, err2.Error
	} else {
		data = item.Data()
		data["createdAt"] = item.CreateTime
		data["id"] = id
	}
	return data, err, err2.Success
}

func Add(collection string, data interface{}) map[string]interface{} {
	documentRef, _, err := client.Collection(collection).Add(ctx, data)
	if err != nil {
		log.Logger.Fatal("Failed adding alert to firebase : %v", err)
	}
	savedData, _, _ := Get(collection, documentRef.ID)
	return savedData
}

func Delete(collection string, id string) (error, int) {
	_, err := client.Collection(collection).Doc(id).Delete(ctx)
	if err != nil {
		log.Logger.Fatal("Failed deleting alert from firebase : %v", err)
		_, _, errCode := Get(collection, id)
		return err, errCode
	}
	return nil, 0
}

func GetAllDocuments() []model.RuleStruct {
	rules := make([]model.RuleStruct, 0)
	if client != nil {
		collectionRef, _ := client.Collections(ctx).GetAll()
		for _, ref := range collectionRef {
			rules = append(rules, model.RuleStruct{
				Owner:  model.RuleOwner{Email: ref.ID},
				Alerts: GetAll(ref.ID),
			})
		}
	} else {
		log.Logger.Error("Firebase client not initialised yet!")
	}
	return rules
}
