package worker

import (
	"canny/pkg/alphavantage"
	"canny/pkg/cache"
	"canny/pkg/config"
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
	return config.Cfg().String("app.currencies.exchange")
}

func getEligibleCoins() []string {
	return config.Cfg().Strings("app.currencies.allowed")
}

func RefreshCache() {
	coins := getEligibleCoins()
	exchange := getEligibleExchangeCurrency()

	//FIXME : the order is still sequential, make it async
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
