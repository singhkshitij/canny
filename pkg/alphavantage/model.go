package alphavantage

type PriceDataResult struct {
	AOpenINR     string `json:"1a. open (INR)"`
	BOpenUSD     string `json:"1b. open (USD)"`
	AHighINR     string `json:"2a. high (INR)"`
	BHighUSD     string `json:"2b. high (USD)"`
	ALowINR      string `json:"3a. low (INR)"`
	BLowUSD      string `json:"3b. low (USD)"`
	ACloseINR    string `json:"4a. close (INR)"`
	BCloseUSD    string `json:"4b. close (USD)"`
	Volume       string `json:"5. volume"`
	MarketCapUSD string `json:"6. market cap (USD)"`
}

type DailyCurrencyDataResult struct {
	MetaData struct {
		Information         string `json:"1. Information"`
		DigitalCurrencyCode string `json:"2. Digital Currency Code"`
		DigitalCurrencyName string `json:"3. Digital Currency Name"`
		MarketCode          string `json:"4. Market Code"`
		MarketName          string `json:"5. Market Name"`
		LastRefreshed       string `json:"6. Last Refreshed"`
		TimeZone            string `json:"7. Time Zone"`
	} `json:"Meta Data"`
	TimeSeriesDigitalCurrencyDaily map[string]*PriceDataResult `json:"Time Series (Digital Currency Daily)"`
}
