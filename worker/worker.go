package worker

import (
	"canny/pkg/alphavantage"
	"canny/pkg/cache"
	"canny/pkg/log"
	"canny/pkg/scheduler"
)

func Setup() {
	alphavantage.Setup()
	cache.Setup()
	scheduler.Setup()
}

func getEligibleExchangeCurrency() string {
	// TODO get from configs
	return "INR"
}

func getEligibleCoins() []string {
	// TODO use configs to fetch this
	return []string{"BTC", "ETH", "MATIC", "ADA"}
}

func RefreshCache() {
	coins := getEligibleCoins()
	exchange := getEligibleExchangeCurrency()
	for _, coinName := range coins {
		data := alphavantage.GetCurrencyData(coinName, exchange)
		cache.Set(coinName, data)
		log.Logger.Infof("Refreshed cache for coin %s", coinName)
	}
}

func InitialiseData() {
	scheduler.Add(5, RefreshCache)
}
