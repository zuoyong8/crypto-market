package model

//
type CurrencyRateReq struct {
	From string `json:"from" form:"from"`
	To   string `json:"to" form:"to"`
}
