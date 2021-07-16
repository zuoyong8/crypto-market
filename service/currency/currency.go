package currency

import (
	"crypto-market/common/usd"
	"crypto-market/model"
)

//获取汇率
func GetCurrencyRate(req *model.CurrencyRateReq) float64 {
	//
	currency := usd.NewCurrencyRate(req.From, req.To)
	return currency.GetCurrencyRate()
}
