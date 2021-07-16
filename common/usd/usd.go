package usd

import (
	"crypto-market/common/xfloat64"
)

type currency interface {
	//获取汇率
	GetCurrencyRate(string, string) float64
}

//
type CurrencyRate struct {
	//
	currencyRate currency
	//
	from string
	//
	to string
}

//
func NewCurrencyRate(from, to string) *CurrencyRate {
	return &CurrencyRate{
		from:         from,
		to:           to,
		currencyRate: newSinaCurrency(),
	}
}

//获取汇率
func (c *CurrencyRate) GetCurrencyRate() float64 {
	rate := c.currencyRate.GetCurrencyRate(c.from, c.to)

	//比较
	if xfloat64.FromFloatCmp(rate, 0.0) == 0 {
		//
		c.currencyRate = newIt120Currency()
		rate = c.currencyRate.GetCurrencyRate(c.from, c.to)
	}
	return rate

}
