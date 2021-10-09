package model

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

type CreateAlertRequest struct {
	Name       string `json:"name" validate:"required"`
	Property   string `json:"property" validate:"required"`
	Symbol     string `json:"symbol" validate:"required"`
	Operator   string `json:"operator" validate:"required"`
	Value      int64  `json:"value" validate:"required_without=Percentage"`
	Percentage int64  `json:"percentage" validate:"required_without=Value"`
	Currency   string `json:"currency" validate:"required"`
}

type CreateAlertResponse struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Property   string `json:"property"`
	Symbol     string `json:"symbol"`
	Operator   string `json:"operator"`
	Value      int64  `json:"value"`
	Percentage int64  `json:"percentage"`
	Currency   string `json:"currency"`
}
