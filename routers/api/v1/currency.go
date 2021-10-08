package v1

import (
	"canny/pkg/app"
	"canny/pkg/cache"
	"canny/pkg/config"
	"canny/pkg/err"
	"canny/pkg/utils"
	"github.com/gin-gonic/gin"
)

// @Summary Get supported currencies
// @Produce  json
// @Success 200 {object} app.CurrenciesResponse
// @Failure 500 {object} app.Response
// @Router /api/v1/currencies [get]
// @tags currency
func Currencies(c *gin.Context) {
	appG := app.Gin{C: c}
	supportedCoins := config.Cfg().Strings("app.currencies.allowed")
	appG.Response(200, err.Success, supportedCoins)
}

// @Summary Get coin data
// @Produce json
// @Param currency path string true "Symbol"
// @Success 200 {object} app.CurrencyDataResponse
// @Failure 500 {object} app.Response
// @Router /api/v1/currencies/{currency} [get]
// @tags currency
func CurrencyData(c *gin.Context) {
	appG := app.Gin{C: c}
	coinName := c.Param("currency")
	data := cache.Get(coinName)
	if data == nil {
		appG.Response(404, err.NotFound, map[string]string{})
	} else {
		appG.Response(200, err.Success, data)
	}

}

// @Summary Get all coin last price
// @Produce json
// @Success 200 {object} app.AllCurrencyPriceResponse
// @Failure 500 {object} app.Response
// @Router /api/v1/currencies/price [get]
// @tags currency
func AllCurrencyData(c *gin.Context) {
	appG := app.Gin{C: c}
	data := cache.Get(utils.AllCoinPriceKey)
	if data == nil {
		appG.Response(404, err.NotFound, map[string]string{})
	} else {
		appG.Response(200, err.Success, data)
	}
}
