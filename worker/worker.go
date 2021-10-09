package worker

import (
	"canny/domain"
	"canny/model"
	"canny/pkg/alphavantage"
	"canny/pkg/cache"
	"canny/pkg/http"
	"canny/pkg/log"
	"canny/pkg/scheduler"
	"canny/pkg/utils"
	"strings"
	"sync"
)

func Setup() {
	http.Setup()
	cache.Setup()
	scheduler.Setup()
}


func RefreshCache() {

	coins := domain.GetAllSupportedCoins()
	exchange := domain.GetEligibleExchangeCurrency()

	var wg sync.WaitGroup

	allCoinPrices := make(map[string]model.ClosingPrice)

	for _, coinName := range coins {
		wg.Add(1)
		go func(c string, e string) {
			data := alphavantage.GetCurrencyData(c, e)
			cache.Set(c, data)
			allCoinPrices[c] = addClosePriceToAllCoinPrices(data)
			log.Logger.Infof("Refreshed cache for coin: %s", c)
			wg.Done()
		}(coinName, exchange)
	}
	wg.Wait()
	cache.Set(utils.AllCoinPriceKey, allCoinPrices)
	log.Logger.Infof("Finished updating cache for all coins!")
}

func addClosePriceToAllCoinPrices(data *alphavantage.DailyCurrencyDataResponse) model.ClosingPrice {
	return model.ClosingPrice{
		Inr: data.TimeSeriesDigitalCurrencyDaily[strings.Split(data.MetaData.LastRefreshed, " ")[0]].CloseINR,
		Usd: data.TimeSeriesDigitalCurrencyDaily[strings.Split(data.MetaData.LastRefreshed, " ")[0]].CloseUSD,
	}
}

func InitialiseData() {
	scheduler.Add(5, RefreshCache)
}
