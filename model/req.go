package model

//
type CurrencyRateReq struct {
	From string `json:"from" form:"from"`
	To   string `json:"to" form:"to"`
}

//
type IndicatorReq struct {
	Name string `json:"name" form:"name"`
}
