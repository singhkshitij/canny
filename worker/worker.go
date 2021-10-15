package worker

import (
	"canny/domain"
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

	allCoinPrices := make(map[string]alphavantage.LatestPrice)

	for _, coinName := range coins {
		wg.Add(1)
		go func(c string, e string) {
			data := alphavantage.GetCurrencyData(c, e)
			latestPrice := addLatestPriceToAllCoinPrices(data)
			data.MetaData.LatestPrice = latestPrice
			cache.Set(c, data)
			allCoinPrices[c] = latestPrice
			log.Logger.Infof("Refreshed cache for coin: %s", c)
			wg.Done()
		}(coinName, exchange)
	}
	wg.Wait()
	cache.Set(utils.AllCoinPriceKey, allCoinPrices)
	log.Logger.Infof("Finished updating cache for all coins!")
}

func addLatestPriceToAllCoinPrices(data *alphavantage.DailyCurrencyDataResponse) alphavantage.LatestPrice {

	latestData := data.TimeSeriesDigitalCurrencyDaily[strings.Split(data.MetaData.LastRefreshed, " ")[0]]

	return alphavantage.LatestPrice{
		OpenINR:  latestData.OpenINR,
		OpenUSD:  latestData.OpenUSD,
		HighINR:  latestData.HighINR,
		HighUSD:  latestData.HighUSD,
		LowINR:   latestData.LowINR,
		LowUSD:   latestData.LowUSD,
		CloseINR: latestData.CloseINR,
		CloseUSD: latestData.CloseUSD,
	}
}

func InitialiseData() {
	scheduler.Add(5, RefreshCache)
}
