package worker

import "canny/pkg/alphavantage"

func Setup() {
	alphavantage.Setup()
}

func initialiseData() {
	// TODO remove this
	alphavantage.GetCurrencyData("BTC", "INR")
}
