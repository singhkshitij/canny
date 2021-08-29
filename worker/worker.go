package worker

import (
	"canny/pkg/alphavantage"
	"canny/pkg/cache"
	"canny/pkg/config"
	"canny/pkg/http"
	"canny/pkg/log"
	"canny/pkg/scheduler"
	"sync"
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

	var wg sync.WaitGroup

	for _, coinName := range coins {
		wg.Add(1)
		go func(c string, e string) {
			data := alphavantage.GetCurrencyData(c, e)
			cache.Set(c, data)
			log.Logger.Infof("Refreshed cache for coin: %s", c)
			wg.Done()
		}(coinName, exchange)
	}

	wg.Wait()
	log.Logger.Infof("Finished updating cache for all coins!")
}

func InitialiseData() {
	scheduler.Add(5, RefreshCache)
}
