package v1

import (
	"canny/domain"
	"canny/model"
	"canny/pkg/err"
	"github.com/gin-gonic/gin"
)

// @Summary Get supported currencies
// @Produce  json
// @Success 200 {object} model.CurrenciesResponse
// @Failure 500 {object} model.Response
// @Router /api/v1/currencies [get]
// @tags currency
func Currencies(c *gin.Context) {
	appG := model.Gin{C: c}
	appG.Response(200, err.Success, domain.GetAllSupportedCoins())
}

// @Summary Get coin data
// @Produce json
// @Param currency path string true "Symbol"
// @Success 200 {object} model.CurrencyDataResponse
// @Failure 500 {object} model.Response
// @Router /api/v1/currencies/{currency} [get]
// @tags currency
func CurrencyData(c *gin.Context) {
	appG := model.Gin{C: c}
	coinName := c.Param("currency")
	data := domain.GetCoinCurrencyData(coinName)
	if data == nil {
		appG.Response(404, err.NotFound, map[string]string{})
	} else {
		appG.Response(200, err.Success, data)
	}

}

// @Summary Get all coin last price
// @Produce json
// @Success 200 {object} model.AllCurrencyPriceResponse
// @Failure 500 {object} model.Response
// @Router /api/v1/currencies/price [get]
// @tags currency
func AllCurrencyData(c *gin.Context) {
	appG := model.Gin{C: c}
	data := domain.GetAllCurrencyData()
	if data == nil {
		appG.Response(404, err.NotFound, map[string]string{})
	} else {
		appG.Response(200, err.Success, data)
	}
}
