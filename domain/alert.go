package domain

import (
	"canny/model"
	"canny/pkg/firebase"
)

var sampleEmail = "kshitijzxjava@gmail.com"

func CreateAlert(data model.CreateAlertRequest) (map[string]interface{}, error) {
	//Ideally this would be extracted from bearer token
	savedData := firebase.Add(sampleEmail, data)
	return savedData, nil
}

func GetAllAlerts() []map[string]interface{} {
	return firebase.GetAll(sampleEmail)
}
