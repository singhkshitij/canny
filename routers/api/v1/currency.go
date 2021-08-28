package v1

import (
	"canny/pkg/app"
	"canny/pkg/cache"
	"canny/pkg/err"
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
	// TODO use configs to fetch this
	supportedCoins := []string{"BTC", "ETH", "MATIC"}
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
