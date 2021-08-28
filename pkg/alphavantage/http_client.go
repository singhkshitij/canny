package alphavantage

import (
	"canny/pkg/log"
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

var client *resty.Client

func Setup() {
	client = resty.New()
}

func GetCurrencyData(symbol string, market string) *DailyCurrencyDataResult {
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

func handleResponse(resp *resty.Response) *DailyCurrencyDataResult {

	parsedResponse := &DailyCurrencyDataResult{}
	statusCode := resp.StatusCode()

	if statusCode >= 200 && statusCode < 300 {
		json.Unmarshal(resp.Body(), &parsedResponse)
		return parsedResponse
	} else {
		log.Logger.Errorf("Call failed and server returned error code %d", statusCode)
	}
	return nil
}
