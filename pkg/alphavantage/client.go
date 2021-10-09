package alphavantage

import (
	"canny/pkg/config"
	"canny/pkg/http"
	"canny/pkg/log"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"sort"
)

func GetCurrencyData(symbol string, market string) *DailyCurrencyDataResponse {
	log.Logger.Debugf("Getting data for currency %s", symbol)

	url := config.Cfg().String("alphavantage.client.url")
	params := map[string]string{
		"function": config.Cfg().String("alphavantage.client.fn"),
		"symbol":   symbol,
		"market":   market,
		"apikey":   config.Cfg().String("alphavantage.client.apiKey"),
	}

	resp, err := http.Get(url, params)
	if err != nil {
		log.Logger.Errorf("Failed to get currency data for currency %s, recieved response code %s", symbol, err.Error())
		return nil
	}
	return handleResponse(resp)
}

func handleResponse(resp *resty.Response) *DailyCurrencyDataResponse {

	parsedResponse := &DailyCurrencyDataResponse{}
	statusCode := resp.StatusCode()

	if statusCode >= 200 && statusCode < 300 {
		json.Unmarshal(resp.Body(), &parsedResponse)
		limitedTimeSeriesDigitalCurrencyDaily := limitNumberOfEntriesAndTransformKeys(parsedResponse)
		parsedResponse.TimeSeriesDigitalCurrencyDaily = limitedTimeSeriesDigitalCurrencyDaily
		return parsedResponse
	} else {
		log.Logger.Errorf("Call failed and server returned error code %d", statusCode)
	}
	return nil
}

func limitNumberOfEntriesAndTransformKeys(response *DailyCurrencyDataResponse) map[string]*PriceDataResponse {
	timeSeries := response.TimeSeriesDigitalCurrencyDaily
	keys := make([]string, 0, len(timeSeries))
	for key := range timeSeries {
		keys = append(keys, key)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	limitedTimeSeriesDates := make([]string, config.Cfg().Int("model.limit.data.currency"))
	copy(limitedTimeSeriesDates, keys)
	limitedTimeSeriesData := make(map[string]*PriceDataResponse)

	for _, date := range limitedTimeSeriesDates {
		limitedTimeSeriesData[date] = timeSeries[date]
	}

	return limitedTimeSeriesData
}
