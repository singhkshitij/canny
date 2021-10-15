package alphavantage

import (
	"encoding/json"
)

type PriceDataResponse struct {
	OpenINR      string `json:"1a. open (INR)"`
	OpenUSD      string `json:"1b. open (USD)"`
	HighINR      string `json:"2a. high (INR)"`
	HighUSD      string `json:"2b. high (USD)"`
	LowINR       string `json:"3a. low (INR)"`
	LowUSD       string `json:"3b. low (USD)"`
	CloseINR     string `json:"4a. close (INR)"`
	CloseUSD     string `json:"4b. close (USD)"`
	Volume       string `json:"5. volume"`
	MarketCapUSD string `json:"6. market cap (USD)"`
}

type LatestPrice struct {
	OpenINR  string `json:"openINR"`
	OpenUSD  string `json:"openUSD"`
	HighINR  string `json:"highINR"`
	HighUSD  string `json:"highUSD"`
	LowINR   string `json:"lowINR"`
	LowUSD   string `json:"lowUSD"`
	CloseINR string `json:"closeINR"`
	CloseUSD string `json:"closeUSD"`
}

type MetaDataResponse struct {
	Information         string      `json:"1. Information"`
	DigitalCurrencyCode string      `json:"2. Digital Currency Code"`
	DigitalCurrencyName string      `json:"3. Digital Currency Name"`
	MarketCode          string      `json:"4. Market Code"`
	MarketName          string      `json:"5. Market Name"`
	LastRefreshed       string      `json:"6. Last Refreshed"`
	TimeZone            string      `json:"7. Time Zone"`
	LatestPrice         LatestPrice `json:"8. Latest Price"`
}

type DailyCurrencyDataResponse struct {
	MetaData                       MetaDataResponse              `json:"Meta Data"`
	TimeSeriesDigitalCurrencyDaily map[string]*PriceDataResponse `json:"Time Series (Digital Currency Daily)"`
}

type PriceDataObj struct {
	OpenINR      string `json:"openINR"`
	OpenUSD      string `json:"openUSD"`
	HighINR      string `json:"highINR"`
	HighUSD      string `json:"highUSD"`
	LowINR       string `json:"lowINR"`
	LowUSD       string `json:"lowUSD"`
	CloseINR     string `json:"closeINR"`
	CloseUSD     string `json:"closeUSD"`
	Volume       string `json:"volume"`
	MarketCapUSD string `json:"marketCapUSD"`
}

type MetaDataObj struct {
	Information         string      `json:"information"`
	DigitalCurrencyCode string      `json:"digitalCurrencyCode"`
	DigitalCurrencyName string      `json:"digitalCurrencyName"`
	MarketCode          string      `json:"marketCode"`
	MarketName          string      `json:"marketName"`
	LastRefreshed       string      `json:"lastRefreshed"`
	TimeZone            string      `json:"timeZone"`
	LatestPrice         LatestPrice `json:"latestPrice"`
}

type DailyCurrencyDataObj struct {
	MetaData                       MetaDataResponse              `json:"metaData"`
	TimeSeriesDigitalCurrencyDaily map[string]*PriceDataResponse `json:"timeSeries"`
}

func (u *MetaDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(MetaDataObj(*u))
}

func (u *PriceDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(PriceDataObj(*u))
}

func (u *DailyCurrencyDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(DailyCurrencyDataObj(*u))
}
