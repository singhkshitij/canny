package meta

import (
	"canny/pkg/alphavantage"
	"canny/pkg/err"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode,
		Msg:  err.GetMsg(errCode),
		Data: data,
	})
	return
}

type CurrenciesResponse struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data []string `json:"data"`
}

type CurrencyDataResponse struct {
	Code int                                    `json:"code"`
	Msg  string                                 `json:"msg"`
	Data alphavantage.DailyCurrencyDataResponse `json:"data"`
}

type ClosingPrice struct {
	Inr string `json:"inr"`
	Usd string `json:"usd"`
}

type AllCurrencyPriceResponse struct {
	Data map[string]ClosingPrice `json:"data"`
}