package worker

import (
	"canny/pkg/alphavantage"
	"canny/pkg/cache"
	"canny/pkg/http"
	"canny/pkg/log"
	"canny/pkg/scheduler"
)

func Setup() {
	http.Setup()
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
		responseChannel := make(chan *alphavantage.DailyCurrencyDataResponse)
		go alphavantage.GetCurrencyData(coinName, exchange, responseChannel)
		data := <-responseChannel
		cache.Set(coinName, data)
		log.Logger.Infof("Refreshed cache for coin %s", coinName)
	}
}

func InitialiseData() {
	scheduler.Add(5, RefreshCache)
}
