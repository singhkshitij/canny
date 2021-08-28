package v1

import (
	"canny/pkg/app"
	"canny/pkg/err"
	"github.com/gin-gonic/gin"
)

// @Summary Get supported currencies
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/currencies [get]
// @tags currency
func Currencies(c *gin.Context) {
	appG := app.Gin{C: c}
	supportedCoins := []string{"BTC", "ETH", "MATIC"}
	appG.Response(200, err.SUCCESS, supportedCoins)
}
