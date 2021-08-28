package v1

import (
	"canny/pkg/app"
	"canny/pkg/err"
	"github.com/gin-gonic/gin"
)

func Currencies(c *gin.Context) {
	appG := app.Gin{C: c}
	supportedCoins := []string{"BTC", "ETH", "MATIC"}
	appG.Response(200, err.SUCCESS, supportedCoins)
}
