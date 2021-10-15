package model

import (
	"canny/pkg/alphavantage"
	"canny/pkg/err"
	"github.com/gin-gonic/gin"
	"time"
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

type AllCurrencyPriceResponse struct {
	Code int                                 `json:"code"`
	Msg  string                              `json:"msg"`
	Data map[string]alphavantage.LatestPrice `json:"data"`
}

type CreateAlertRequest struct {
	Name       string  `json:"name" validate:"required"`
	Property   string  `json:"property" validate:"required"`
	Operator   string  `json:"operator" validate:"required"`
	Value      float64 `json:"value" validate:"required_without=Percentage,gte=0"`
	Percentage int64   `json:"percentage" validate:"required_without=Value,gte=0,lte=100"`
	Currency   string  `json:"currency" validate:"required"`
}

type CreateAlertResponseData struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Property   string    `json:"property"`
	Operator   string    `json:"operator"`
	Value      float64     `json:"value"`
	Percentage int64     `json:"percentage"`
	Currency   string    `json:"currency"`
	CreatedAt  time.Time `json:"createdAt"`
}

type CreateAlertResponse struct {
	Code int                     `json:"code"`
	Msg  string                  `json:"msg"`
	Data CreateAlertResponseData `json:"data"`
}

type GetAllAlertResponse struct {
	Code int                       `json:"code"`
	Msg  string                    `json:"msg"`
	Data []CreateAlertResponseData `json:"data"`
}

type DryRunAlertStatus struct {
	Passed bool `json:"passed"`
}

type DryRunAlertResponse struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data DryRunAlertStatus `json:"data"`
}

type RulePipelineDataStruct struct {
	Data             CreateAlertResponseData
	CoinCurrentPrice float64
}

type RuleOwner struct {
	Email string `json:"email"`
}

type RuleStruct struct {
	Owner  RuleOwner                 `json:"owner"`
	Alerts []CreateAlertResponseData `json:"data"`
}
