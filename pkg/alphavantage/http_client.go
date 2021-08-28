package alphavantage

import (
	"canny/pkg/log"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"sort"
)

var client *resty.Client

func Setup() {
	client = resty.New()
}

// TODO make this an async process so that all coin calls can be in parellel
func GetCurrencyData(symbol string, market string) *DailyCurrencyDataResponse {

	// TODO move api key to config
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"function": "DIGITAL_CURRENCY_DAILY",
			"symbol":   symbol,
			"market":   market,
			"apikey":   "KIKJ1AN4SPAXV1BO",
		}).
		SetHeader("Accept", "application/json").
		Get("https://www.alphavantage.co/query")

	if err != nil {
		log.Logger.Errorf("Failed to get currency data for %s. Recieved response code %s", symbol, err.Error())
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

	//TODO this value has to come from config
	limitedTimeSeriesDates := make([]string, 365)
	copy(limitedTimeSeriesDates, keys)
	limitedTimeSeriesData := make(map[string]*PriceDataResponse)

	for _, date := range limitedTimeSeriesDates {
		limitedTimeSeriesData[date] = timeSeries[date]
	}

	return limitedTimeSeriesData
}
