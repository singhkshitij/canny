package domain

import (
	"canny/pkg/alphavantage"
	"canny/pkg/utils"
)

func GetPropertyOfData(priceData alphavantage.LatestPrice, property string) string {
	switch property {
	case utils.CloseINR:
		return priceData.CloseINR
	case utils.CloseUSD:
		return priceData.CloseUSD
	case utils.HighINR:
		return priceData.HighINR
	case utils.HighUSD:
		return priceData.HighUSD
	case utils.LowINR:
		return priceData.LowINR
	case utils.LowUSD:
		return priceData.LowUSD
	case utils.OpenUSD:
		return priceData.OpenUSD
	case utils.OpenINR:
		return priceData.OpenINR
	default:
		return priceData.CloseUSD
	}
}
