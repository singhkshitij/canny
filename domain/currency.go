package domain

import (
	"canny/pkg/cache"
	"canny/pkg/config"
	"canny/pkg/utils"
)

func GetAllSupportedCoins() []string {
	return config.Cfg().Strings("app.currencies.allowed")
}

func GetCoinCurrencyData(coinName string) interface{} {
	return cache.Get(coinName)
}

func GetAllCurrencyData() interface{} {
	return cache.Get(utils.AllCoinPriceKey)
}

func GetEligibleExchangeCurrency() string {
	return config.Cfg().String("app.currencies.exchange")
}
