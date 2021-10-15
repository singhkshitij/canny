package domain

import (
	"canny/pkg/alphavantage"
	"canny/pkg/cache"
	"canny/pkg/config"
	"canny/pkg/err"
	"canny/pkg/utils"
)

func GetAllSupportedCoins() []string {
	return config.Cfg().Strings("app.currencies.allowed")
}

func GetCoinCurrencyData(coinName string) interface{} {
	return cache.Get(coinName)
}

func GetAllCurrencyData() (map[string]alphavantage.LatestPrice, int) {
	result := cache.Get(utils.AllCoinPriceKey)
	if result != nil {
		return result.(map[string]alphavantage.LatestPrice), 0
	}
	return map[string]alphavantage.LatestPrice{}, err.NotFound
}

func GetEligibleExchangeCurrency() string {
	return config.Cfg().String("app.currencies.exchange")
}
