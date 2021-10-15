package domain

import (
	"canny/model"
	"canny/pkg/err"
	"canny/pkg/firebase"
	"canny/pkg/utils"
	"strconv"
	"time"
)

//Ideally this would be extracted from bearer token
var sampleEmail = "kshitijzxjava@gmail.com"

func CreateAlert(data model.CreateAlertRequest) (map[string]interface{}, error) {

	savedData := firebase.Add(sampleEmail, data)
	return savedData, nil
}

func GetAllAlerts() []model.CreateAlertResponseData {
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

func DryRunAlert(data model.CreateAlertRequest) bool {
	return EvaluateRule(model.CreateAlertResponseData{
		Id:         "",
		Name:       data.Name,
		Property:   data.Property,
		Operator:   data.Operator,
		Value:      data.Value,
		Percentage: data.Percentage,
		Currency:   data.Currency,
		CreatedAt:  time.Time{},
	})
}

func EvaluateRule(data model.CreateAlertResponseData) bool {

	pipelineData := model.RulePipelineDataStruct{Data: data}
	pipelineData.CoinCurrentPrice, _ = GetCurrentPriceOfCoin(data.Currency, data.Property)

	if data.Percentage != 0 {
		data.Value = getPercentageValue(data.Percentage, data.Operator, pipelineData.CoinCurrentPrice)
	}

	switch data.Operator {
	case utils.EqualsTo:
		return data.Value == pipelineData.CoinCurrentPrice
	case utils.LessThan:
		return data.Value > pipelineData.CoinCurrentPrice
	case utils.LessThanEqualTo:
		return data.Value >= pipelineData.CoinCurrentPrice
	case utils.GreaterThan:
		return data.Value < pipelineData.CoinCurrentPrice
	case utils.GreaterThanEqualTo:
		return data.Value <= pipelineData.CoinCurrentPrice
	default:
		return false
	}
}

func GetCurrentPriceOfCoin(currency string, property string) (float64, error) {
	data, _ := GetAllCurrencyData()
	strValue := GetPropertyOfData(data[currency], property)
	return strconv.ParseFloat(strValue, 64)
}

func getPercentageValue(percentage int64, operator string, price float64) float64 {
	percentageValue := price * float64(percentage/100)
	switch operator {
	case utils.GreaterThan, utils.GreaterThanEqualTo:
		return price + percentageValue
	case utils.LessThan, utils.LessThanEqualTo:
		return price - percentageValue
	default:
		return price
	}
}
