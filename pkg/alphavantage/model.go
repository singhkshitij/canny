package alphavantage

type PriceDataResponse struct {
	CloseINR     string `json:"4a. close (INR)"`
	CloseUSD     string `json:"4b. close (USD)"`
}

type MetaData struct {
	DigitalCurrencyCode string `json:"2. Digital Currency Code"`
	DigitalCurrencyName string `json:"3. Digital Currency Name"`
	LastRefreshed       string `json:"6. Last Refreshed"`
	TimeZone            string `json:"7. Time Zone"`
}

type DailyCurrencyDataResponse struct {
	MetaData                       MetaData                      `json:"Meta Data"`
	TimeSeriesDigitalCurrencyDaily map[string]*PriceDataResponse `json:"Time Series (Digital Currency Daily)"`
}
