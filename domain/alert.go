package domain

import (
	"canny/model"
	"canny/pkg/err"
	"canny/pkg/firebase"
)

//Ideally this would be extracted from bearer token
var sampleEmail = "kshitijzxjava@gmail.com"

func CreateAlert(data model.CreateAlertRequest) (map[string]interface{}, error) {

	savedData := firebase.Add(sampleEmail, data)
	return savedData, nil
}

func GetAllAlerts() []map[string]interface{} {
	return firebase.GetAll(sampleEmail)
}

func DeleteAlert(alertId string) (error, int) {
	er, code := firebase.Delete(sampleEmail, alertId)
	if code == 404 {
		return nil, err.NotFound
	} else if er != nil {
		return er, 0
	}
	return nil, err.Success
}

func GetAlert(alertId string) (map[string]interface{}, error, int) {
	data, er, errCode := firebase.Get(sampleEmail, alertId)
	if errCode == err.NotFound {
		return nil, er, err.NotFound
	} else if er != nil {
		return nil, er, errCode
	}
	return data, nil, err.Success
}
